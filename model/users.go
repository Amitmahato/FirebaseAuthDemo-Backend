package model

type Users struct {
	Id         uint64 `sql:"AUTO_INCREMENT" json:"user_id,omitempty"`
	UID        string `json:"user_uid,omitempty"`
	FirstName  string `json:"user_firstName,omitempty"`
	LastName   string `json:"user_lastName,omitempty"`
	Email      string `json:"user_email" validate:"required"`
	Phone      string `json:"user_phone,omitempty"`
	Created_at string `json:"user_created_at,omitempty"`
	Updated_at string `json:"user_updated_at,omitempty"`
}

// TableName func returns the table name
func (e *Users) TableName() string {
	return "users"
}
