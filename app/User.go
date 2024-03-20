package app

type UserRegister struct {
	Username string `valid:"type(string)"`
	Email    string `valid:"email"`
	Password string `valid:"type(string),minstringlength(6)"`
}

type UserLogin struct {
	Email    string `valid:"email"`
	Password string `valid:"type(string),minstringlength(6)"`
}
