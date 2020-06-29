package main

type Developer struct {

	//gorm.Model

	ID			uint		`json:"id" gorm:"primary_key"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Skills		[]*Skill	`json:"skills" gorm:"many2many:developer_skill"`


}


type Skill struct {
	ID			uint		`json:"id" gorm:"primary_key"`
	Name 		string		`json:"name"`
	Description	string		`json:"description"`
	Developers	[]*Developer `json:"developers" gorm:"many2many:developer_skill"`			

}

// ErrorResponse is interface for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}
