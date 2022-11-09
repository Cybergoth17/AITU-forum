package controllers

type Admin struct {
	IsAdmin bool
	Wrapper Opportunities
	Auth
}

func (a *Admin) DoYourOwnJob() string {
	return a.Wrapper.DoYourOwnJob() + "Moderate posts, users, send notifications."
}
