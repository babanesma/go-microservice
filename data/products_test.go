package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "n",
		Price: 1.00,
		SKU:   "abc-def-ghe",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}