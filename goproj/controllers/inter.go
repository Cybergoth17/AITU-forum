package controllers

type Noauth struct{}

func (n *Noauth) DoYourOwnJob() string {
	return "Have access to all questions."
}
