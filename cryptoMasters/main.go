package main

import "leo/cryptomasters/api"

func main() {
	rate, err := api.GetRate("BTC")
	print(rate, err)
}
