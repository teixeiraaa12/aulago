package main

import "fmt"

func main(){
menu := [string]float64{
"pizza": 40.00,
"suco":  12.50,
"xtudo": 20.00,
}
fmt.Println(menu)

for k,v:= range menu{
fmt.Println(k, "R$",v)
}
for key, value := range menu {
fmt.Println(key, "R$", value)
}
