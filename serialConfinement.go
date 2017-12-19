package cake

// Cake demo
type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touch this cake again
	}

}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for c := range cooked {
		c.state = "iced"
		iced <- c // icer  never touch this cake again
	}
}
