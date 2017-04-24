// Package cake provides serial confinement
package cake

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake //baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooker <-chan *Cakeparams) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
	}

}
