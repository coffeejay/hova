package main 

import "hova_dir/hova"

/*  
  
  Add routes to hova with the 
  corresponding view and controller action
  name. 

  Initialize hova when all routes have been added.

  ex. 	
  	hova.AddRoute("/","index.html","Index")
  	hova.Init()

  The controller action must be defined in controller.go
  with the same name. (Use proper Go namespacing)

*/


func main(){

	hova.Init()

}
