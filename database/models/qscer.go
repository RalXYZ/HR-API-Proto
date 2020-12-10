package models

type Qscer struct {
	UID      uint   `gorm:"primary_key;auto_increment"`
	Zjuid    string `gorm:"size:10;not_null"`
	Name     string `gorm:"size:10;not_null"`
	Qscid    string `gorm:"size:10;not_null"`
	Birthday string `gorm:"size:10;not_null"`
}
