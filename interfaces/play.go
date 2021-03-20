package main

type TaskSystem interface {
	calculateTax() int
}

type indiaTax struct {
	taxPercentage int
	income        int
}

func (i *indiaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type nigeiraTax struct {
	taxPercentage int
	income        int
}

func (n *nigeiraTax) calculateTax() int {
	tax := n.income * n.taxPercentage / 200
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (u *usaTax) calculateTax() int {
	tax := u.income * u.taxPercentage / 50
	return tax
}

func calculateTaxes(tasksystems []TaskSystem) int {
	total := 0
	for _, tasksystem := range tasksystems {
		total += tasksystem.calculateTax()
	}
	return total
}
