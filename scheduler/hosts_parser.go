package scheduler

import (
	"bytes"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

type segment interface {
	validate() error
	expand() []string
}

type segments []segment

type segString string

type segRange string

var segStringRegex = regexp.MustCompile(`^[a-zA-Z0-9\-\*\.]*$`)

func (seg segString) validate() error {
	if !segStringRegex.MatchString(string(seg)) {
		return errors.Errorf("segment %s contains invalid characters", seg)
	}
	return nil
}

func (seg segString) expand() []string {
	return []string{string(seg)}
}

var segRangeRegex = regexp.MustCompile(`^\[(\d+|[a-zA-Z]):(\d+|[a-zA-Z])]$`)

func (seg segRange) getRange() (string, string, error) {
	res := segRangeRegex.FindAllStringSubmatch(string(seg), -1)
	if len(res) != 1 {
		return "", "", errors.Errorf("segment %s contains invalid characters or formats", seg)
	}
	return res[0][1], res[0][2], nil
}

var isChars = regexp.MustCompile(`^[a-zA-Z]$`)

func (seg segRange) validate() error {
	start, end, err := seg.getRange()
	if err != nil {
		return err
	}

	if startInt, err := strconv.Atoi(start); err == nil {
		if endInt, err := strconv.Atoi(end); err != nil {
			return errors.Errorf("range start and end must both be digits or non-digit characters")
		} else if startInt > endInt {
			return errors.Errorf("range start must be smaller than end")
		} else if (len(start) != len(strconv.Itoa(startInt))) && (len(start) != len(end)) {
			return errors.Errorf("range start and end must be equal-length'd")
		} else {
			return nil
		}
	}

	if !isChars.MatchString(end) {
		return errors.Errorf("range start and end must both be digits or non-digit characters")
	} else if start > end {
		return errors.Errorf("range start must be alphabetically smaller than end")
	}

	return nil
}

func (seg segRange) expand() []string {
	start, end, _ := seg.getRange()

	var res []string
	if startInt, err := strconv.Atoi(start); err == nil {
		endInt, _ := strconv.Atoi(end)
		for i := startInt; i <= endInt; i++ {
			res = append(res, pad(i, len(start)))
		}
	} else {
		startR := []rune(start)[0]
		endR := []rune(end)[0]
		for startR <= endR {
			res = append(res, string(startR))
			startR++
		}
	}

	return res
}

func pad(n, l int) string {
	var buffer bytes.Buffer
	for i := len(strconv.Itoa(n)); i < l; i++ {
		buffer.WriteString("0")
	}
	buffer.WriteString(strconv.Itoa(n))
	return buffer.String()
}

func (segs *segments) add(seg segment) error {
	if err := seg.validate(); err != nil {
		return err
	}
	*segs = append(*segs, seg)
	return nil
}

func (segs segments) merge() []string {
	var res []string
	for _, seg := range segs {
		res = mergeStrings(res, seg.expand())
	}
	return res
}

func mergeStrings(list1, list2 []string) []string {

	if len(list1) == 0 {
		return list2
	}
	var newList []string
	for _, s1 := range list1 {
		for _, s2 := range list2 {
			newList = append(newList, s1+s2)
		}
	}
	return newList
}

func parseHosts(hosts string) ([]string, error) {
	var buffer bytes.Buffer

	var segs segments

	for _, c := range hosts {
		switch c {
		case '[':
			if err := segs.add(segString(buffer.String())); err != nil {
				return nil, errors.Wrap(err, "malformed host")
			}
			buffer.Reset()
			buffer.WriteRune(c)
		case ']':
			buffer.WriteRune(c)
			if err := segs.add(segRange(buffer.String())); err != nil {
				return nil, errors.Wrap(err, "malformed host")
			}
			buffer.Reset()
		default:
			buffer.WriteRune(c)
		}
	}

	if err := segs.add(segString(buffer.String())); err != nil {
		return nil, errors.Wrap(err, "malformed host")
	}
	return segs.merge(), nil
}

func parseHostsData(m *map[interface{}]interface{}) *map[interface{}]interface{} {
	if m == nil {
		return m
	}
	for key, value := range *m {
		if key == "hosts" {
			var list []string
			original, ok := value.(string)
			if ok {
				for _, h := range strings.Split(original, " ") {
					l, _ := parseHosts(h)
					list = append(list, l...)
				}
			} else {
				original, ok := value.(map[interface{}]interface{})
				if ok {
					for innerkey := range original {
						original, ok := innerkey.(string)
						if ok {
							l, _ := parseHosts(original)
							list = append(list, l...)
						}
					}
				}

			}
			(*m)[key] = list
		} else {
			original, ok := value.(map[interface{}]interface{})
			if ok {
				value = parseHostsData(&original)
			}
		}
	}

	return m
}
