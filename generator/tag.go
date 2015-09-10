package generator

import (
	"strings"
)

type TagField struct {
	TagKey   string
	TagValue string
	IsKV     bool // if true TagFile 形如 foo=bar,if false ,only TagKey 有用
}
type TagFieldList []TagField

func (l TagFieldList) Contains(field string) bool {
	for _, tagField := range l {
		if tagField.TagKey == field {
			return true
		}
	}
	return false
}
func (l TagFieldList) GetValue(tagKey string) string {
	for _, tagField := range l {
		if tagField.TagKey == tagKey && tagField.IsKV {
			return tagField.TagValue
		}
	}
	return ""
}

// 比如`mysql:"type=int,pk"` 分别对应 TagField{TagKey:"type",TagValue:"int",IsKV:true} {TagKey:"pk"}

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (tagFields []TagField) {
	for _, tagFieldStr := range strings.Split(tag, ",") {
		i := strings.Index(tagFieldStr, "=")
		if i != -1 {
			tagField := TagField{
				IsKV:     true,
				TagKey:   tagFieldStr[:i],
				TagValue: tagFieldStr[i+1:],
			}
			tagFields = append(tagFields, tagField)
		} else {
			tagField := TagField{
				IsKV:   false,
				TagKey: tagFieldStr,
			}
			tagFields = append(tagFields, tagField)

		}

	}
	return
}
