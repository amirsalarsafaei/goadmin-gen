package fields

type FieldType string

var (
	Number FieldType = "Number"
	String FieldType = "String"
	Image  FieldType = "Image"
)

func FieldTypeToDBType(f FieldType) string {
	switch f {
	case Number:
		return "Int"
	case String, Image:
		return "String"
	default:
		return ""
	}
}
