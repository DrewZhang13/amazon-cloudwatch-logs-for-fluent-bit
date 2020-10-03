package cloudwatch

import (
	"io"
	"strconv"
	"strings"

	"github.com/valyala/fasttemplate"
)

// tagKeysToMap converts a raw string into a go map.
// This is used by input data to create AWS tags applied to newly-created log groups.
//
// The input string should be match this: "key=value,key2=value2".
// Spaces are trimmed, empty values are permitted, empty keys are ignored.
// The final value in the input string wins in case of duplicate keys.
func tagKeysToMap(tags string) map[string]*string {
	output := make(map[string]*string)

	for _, tag := range strings.Split(strings.TrimSpace(tags), ",") {
		split := strings.SplitN(tag, "=", 2)
		key := strings.TrimSpace(split[0])
		value := ""

		if key == "" {
			continue
		}

		if len(split) > 1 {
			value = strings.TrimSpace(split[1])
		}

		output[key] = &value
	}

	if len(output) == 0 {
		return nil
	}

	return output
}

// parseKeysTemplate takes in an interface map and a list of nested keys. It returns
// the value of the final key, or the name of the first key not found in the chain.
// example keys := "['level1']['level2']['level3']"
// This is called by parseDataMapTags any time a nested value is found in a log Event.
// This procedure checks if any of the nested values match variable identifiers in the logStream or logGroups.
func parseKeysTemplate(data map[interface{}]interface{}, keys string, sanitize func(b []byte) []byte) (string, error) {
	return fasttemplate.ExecuteFuncStringWithErr(keys, "['", "']",
		func(w io.Writer, tag string) (int, error) {
			switch val := data[tag].(type) {
			case []byte:
				return w.Write(sanitize(val))
			case string:
				return w.Write(sanitize([]byte(val)))
			case map[interface{}]interface{}:
				data = val // drill down another level.
				return 0, nil
			default: // missing
				return w.Write(sanitize([]byte(tag)))
			}
		})
}

// parseDataMapTags parses the provided tag values in template form,
// from an interface{} map (expected to contain strings or more interface{} maps).
// This runs once for every log line.
// Used to fill in any template variables that may exist in the logStream or logGroup names.
// If a sanitize function is provided, it runs before writing the data.
func parseDataMapTags(e *Event, logTags []string, t *fasttemplate.Template,
	sanitize func(b []byte) []byte) (string, error) {

	return t.ExecuteFuncStringWithErr(func(w io.Writer, tag string) (int, error) {
		v := strings.Index(tag, "[")
		if v == -1 {
			v = len(tag)
		}

		if tag[:v] == "tag" {
			switch {
			default: // input string is either `tag` or `tag[`, so return the $tag.
				return w.Write(sanitize([]byte(e.Tag)))
			case len(tag) >= 5: // input string is at least "tag[x" where x is hopefully an integer 0-9.
				// The index value is always in the same position: 4:5 (this is why supporting more than 0-9 is rough)
				if v, _ = strconv.Atoi(tag[4:5]); len(logTags) <= v {
					// not enough dots the tag to satisfy the index position, so return whatever the input string was.
					return w.Write(sanitize([]byte("tag" + tag[4:5])))
				}

				return w.Write(sanitize([]byte(logTags[v])))
			}
		}

		switch val := e.Record[tag[:v]].(type) {
		case string:
			return w.Write(sanitize([]byte(val)))
		case map[interface{}]interface{}:
			keyVal, err := parseKeysTemplate(val, tag[v:], sanitize)
			if err != nil {
				return 0, err
			}

			return w.Write(sanitize([]byte(keyVal)))
		case []byte:
			// we should never land here because the interface{} map should have already been converted to strings.
			return w.Write(sanitize(val))
		default: // missing
			return w.Write(sanitize([]byte(tag)))
		}
	})
}

// sanitizeGroup removes special characters from the log group names bytes.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
func sanitizeGroup(b []byte) []byte {
	for i, r := range b {
		// 45-47 = / . -
		// 48-57 = 0-9
		// 65-90 = A-Z
		// 95 = _
		// 97-122 = a-z
		if r == 95 || (r > 44 && r < 58) ||
			(r > 64 && r < 91) || (r > 96 && r < 123) {
			continue
		}

		b[i] = '_'
	}

	return b
}

// sanitizeStream removes : and * from the log stream bytes.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
func sanitizeStream(b []byte) []byte {
	for i, r := range b {
		if r == '*' || r == ':' {
			b[i] = '_'
		}
	}

	return b
}

func newTemplate(template string) (*fasttemplate.Template, error) {
	return fasttemplate.NewTemplate(template, "$(", ")")
}
