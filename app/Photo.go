package app

type Photo struct {
	Title    string	`valid:"type(string)"`
	Caption  string	`valid:"type(string)"`
}
