package main

func main() {

	message := make(chan <- int)

	message <- 1
	message <- 2

	select {
	case message <- 1:
	default:

	}


	messages := make(<- chan int)

	select {
	case <- messages:
	default:

	}

}
