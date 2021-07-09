package fonts

import "strings"

type FontSet struct {
	Result []*Font
}

func (fs *FontSet) FilterByStyle(token string) *FontSet {
	var result []*Font = make([]*Font, 0)

	for _, f := range fs.Result {
		if strings.Contains(strings.ToLower(f.Style), token) {
			result = append(result, f)
		}
	}

	fs.Result = result

	return fs
}

func (fs *FontSet) FilterByFamily(token string) *FontSet {
	var result []*Font = make([]*Font, 0)

	for _, f := range fs.Result {
		if strings.Contains(strings.ToLower(f.Family), token) {
			result = append(result, f)
		}
	}

	fs.Result = result

	return fs
}
