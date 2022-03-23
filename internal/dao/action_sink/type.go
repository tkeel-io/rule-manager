package action_sink

type Table interface {
	GetName() string
	GetFields() []TableField
}

type TableField interface {
	GetName() string
	GetType() string
	ISPK() bool
}

type BaseFieldType struct {
	Name        string `json:"name"`
	LengthLimit bool   `json:"lengthLimit"`
	MinLength   uint16 `json:"minLength"`
	MaxLength   uint16 `json:"maxLength"`
}
