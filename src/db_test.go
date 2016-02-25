package main

import (
	"testing"
	"fmt"
)

func TestConn(t *testing.T) {
	c := conn()
	if c == nil {
		t.Error("Can not connect")
	}
	c2 := conn()
	if c != c2 {
		t.Error("Should have only one conn")
	}
}

func TestCart(t *testing.T) {
	ClearCart()
	devs0 := ListCart()
	if len(devs0) != 0 {
		t.Error("Bad clear cart")
		fmt.Printf("%+v",devs0)
	}
	dev0 := Developer{Login: "diogok"}
	AddToCart(dev0)
	devs1 := ListCart()
	if len(devs1) != 1 {
		t.Error("Cant insert on cart")
	}
	dev1 := Developer{Login: "diogokid"}
	AddToCart(dev1)
	devs2 := ListCart()
	if len(devs2) != 2 {
		t.Error("Cant insert on cart")
	}

	if devs2[0].Login != "diogok" {
		t.Error("Bad cart listing")
	}
	if devs2[1].Login != "diogokid" {
		t.Error("Bad cart listing")
	}

	DelFromCart("diogokid")
	devs3 := ListCart()
	if len(devs3) != 1 {
		t.Error("Cant del from cart")
	}
}

func TestPurchase(t *testing.T) {
	ClearCart()
	ClearPurchases()
	purchs0 := ListPurchases()
	if len(purchs0) !=0 {
		t.Error("Bad clear purchases")
	}

	dev0 := Developer{Login: "diogok"}
	dev1 := Developer{Login: "diogokid"}
	AddToCart(dev0)
	AddToCart(dev1)

	Purchase()
	cart0 := ListCart()
	if len(cart0) != 0 {
		t.Error("Cart not empty after purchase")
	}

	purchs1 := ListPurchases()
	if len(purchs1) != 1 {
		t.Error("Purchase not saved")
	}
	if purchs1[0][0].Login != "diogok" {
		t.Error("Purchase not saved properly")
	}
}

