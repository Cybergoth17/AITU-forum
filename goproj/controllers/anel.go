package controllers

type Auth struct {
	Name     string
	Age      uint32
	Password string
	Email    string
	Wrapper  Opportunities
}

func (a *Auth) DoYourOwnJob() string {

	return a.Wrapper.DoYourOwnJob() + "Can create new posts,answer to questions, be notified. "
}
