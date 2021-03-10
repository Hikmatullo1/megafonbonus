package billing

var oneMegabaytInDiram int = 6

func Calculate1000(amount int) int {
	var (
		price = 3500
	    res = 0
)
	if(amount - 1000 <= 0){
		return price
	}
	res = ((amount-1000) * oneMegabaytInDiram) + price
	return res
}

func Calculate2000(amount int) int {
	price := 5500
	res := 0
	if(amount - 2000 <= 0){
		return price
	}
	res = ((amount-2000) * oneMegabaytInDiram) + price
	return res
}

func Calculate3000(amount int) int {
	price := 7000
	res := 0
	if(amount - 3000 <= 0){
		return price
	}
	res = ((amount-3000) * oneMegabaytInDiram) + price
	return res
}

func Calculate5000(amount int) int {
	price := 9500
	res := 0
	if(amount - 5000 <= 0){
		return price
	}
	res = ((amount-5000) * oneMegabaytInDiram) + price
	return res
}

func Calculate10000(amount int) int {
	price := 17000
	res := 0
	if(amount - 10000 <= 0){
		return price
	}
	res = ((amount-10000) * oneMegabaytInDiram) + price
	return res
}