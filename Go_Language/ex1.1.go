package main

import (
	"fmt"
)

func main() {
var PriceProduct float64;
var QuantityPurchase float64;

fmt.Print("Enter The price of a product- ");
fmt.Scan(&PriceProduct);

fmt.Print("Enter The Quantity of a product- ");
fmt.Scan(&QuantityPurchase);

FinalPrice:=PriceProduct * QuantityPurchase;

fmt.Println("Total price before tax was " ,float64(FinalPrice));

Gst:=FinalPrice *7/100;

FinalPrice=FinalPrice+Gst;

	fmt.Println("Total price For GSt is " ,float64(Gst));

Qst:=FinalPrice * 7.5/100;
FinalPrice=FinalPrice+Qst;

	fmt.Println("Total price For Qst is " ,float64(Qst));

	fmt.Println("Total price After Both taxes were " ,float64(FinalPrice));

}
