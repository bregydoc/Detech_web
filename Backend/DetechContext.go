package main


type DetechContext struct {
	LoggedUser User
	IsDoctorLogged bool
}

var CurrentContext DetechContext

func GetCurrentContext() *DetechContext {
	return &CurrentContext
}

func (c *DetechContext) GetIdOfCurrentDoctor() string {
	return c.LoggedUser.Id
}

func (c *DetechContext) RefreshDefaultContext(user *User) {
	if user.Id != "" {
		c.LoggedUser.Id = user.Id
		c.IsDoctorLogged = true
	}else{
		c.IsDoctorLogged = false
	}

}
