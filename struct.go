package main
import "fmt"
type animal struct {
	race string 
	patte int
	
}
type chien struct {
	name string
	cris string 
	aliment string
	info string
	animal *animal
	
	
}
type chat struct {
	name string
	cris string 
	aliment string
	info string
	animal *animal
	
}
func main (){
	chat := chat {
		name: "milou",
		cris: "mignole",
		aliment: "poison",
		info: "tres bon",
		animal : &animal {
			race: "chatons",
			patte: 4,
		},
	}
	//pour changer les infortion de notre structure
	// chat.name = "lolo"
	chat.name = "lolo"
	fmt.Println(chat.wouwou())
}
func (chient1 chat) wouwou()  string {
	
 	return chient1.name 

} 
