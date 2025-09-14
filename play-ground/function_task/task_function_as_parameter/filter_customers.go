package taskfunctionasparameter

type Customer struct {
	Name string
	Age  int
}

type Product struct {
	Name     string
	Category string
	Price    int
}

type Filter func(Customer) bool

func FiltersCustomer(customer []Customer, filter Filter) []Customer {
	var result []Customer
	for _, v := range customer {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterProducts(products []Product, filterFunc func(Product) bool) []Product {
	var result []Product
	for _, v := range products {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
	return result
}
