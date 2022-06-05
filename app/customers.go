package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {
	gamiolis1 := Customer{FirstName: "1", LastName: "Gamiolis"}
	customers = append(
		customers,
		gamiolis1,
		Customer{FirstName: "2", LastName: "Gamiolis"},
		Customer{FirstName: "3", LastName: "Gamiolis"},
		Customer{FirstName: "4", LastName: "Gamiolis"},
		Customer{FirstName: "5", LastName: "Gamiolis"},
		Customer{FirstName: "6", LastName: "Gamiolis"},
		Customer{FirstName: "7", LastName: "Gamiolis"},
		Customer{FirstName: "8", LastName: "Gamiolis"},
	)
	return customers
}
