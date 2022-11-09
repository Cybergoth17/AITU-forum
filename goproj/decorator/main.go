package decorator

import (
	"fmt"
	"hello/controllers"
)

func main() {
	n := &controllers.Noauth{}
	a := &controllers.Auth{Name: "Anel", Age: 18, Password: "1234567", Email: "211395@astanait.edu.kz", Wrapper: n}
	admin := &controllers.Admin{IsAdmin: true, Wrapper: a}
	fmt.Println("Now let's check")
	fmt.Printf("No authorized: %s\n", n.DoYourOwnJob())
	fmt.Printf("Authorized: %s\n", a.DoYourOwnJob())
	fmt.Printf("Admin: %s\n", admin.DoYourOwnJob())
}
