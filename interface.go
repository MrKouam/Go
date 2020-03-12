package main 
import (	
		"fmt"
		"math"
)


type rectangle struct {
	longeur int 
	largeur int 
}
type carre struct{
	coter int 
	coter2 int 
	
}
func main () {
	car:= rectangle {
		longeur: 4,
		largeur:4,

	}
	rec := rectangle {
		longeur: 4,
		largeur: 5,
	}
	cir := circle {
		rayon:4,
	}
	fmt.Println("le Perimetre")
	fmt.Println("le perimetre du carré","",car.perimetre())
	fmt.Println("le perimetre du rectangle","",rec.perimetre())
	fmt.Println("l'aire")
	fmt.Println("l'aire du carrée"," ",car.laire())
	fmt.Println("l'aire du rectangle","",rec.laire())
	fmt.Println("le Rayon d'un cercle:","",cir.lai())
}
func (r rectangle) perimetre () int {
	return (r.longeur+ r.largeur) * 2
}
func (b rectangle) laire () int {
	return b.longeur * b.largeur 
}
type circle struct {
	rayon float64
}
func (l circle) lai () float64 {
	return math.Pow(l.rayon,l.rayon) * math.Pi
}
