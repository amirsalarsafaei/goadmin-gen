package infotable

import "goadmin-gen/internal/fields"

const (
	imageDisplay = `.FieldDisplay(func(value types.FieldModel) interface{} {
		avatar, ok := value.Row["avatar"].(string)
		if !ok || avatar == "" {
			return "No image provided!"
		}
		return "<img src=\"" + avatar + "\" style=\"max-width:50px;\"></img>"
	})`
)

func GetFieldDisplayByType(f fields.FieldType) string {
	switch f {
	case fields.Image:
		return imageDisplay
	default:
		return ""
	}
}
