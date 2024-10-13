package Models

type Student struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Marks     string `json:"marks"`
}

func (b *Student) TableName() string {
	return "student"
}
