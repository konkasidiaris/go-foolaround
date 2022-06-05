package main

import (
	"fmt"
	"strconv"
)

func main() {
	customers := GetCustomers()

	for customer := range customers {
		fmt.Println(customer)
	}
}

func getData(customerIds []int) (customers []string) {
	for customerId := range customerIds {
		customers = append(customers, "Gamimenos pelatis "+strconv.Itoa(customerId))
	}
	return customers
}
