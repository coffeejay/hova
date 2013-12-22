package hova 

import (
	"html/template"
	"net/http"
	"fmt"
	"reflect"
	"hova_dir/controllers"
)

var view_name, action_name string 

func AddRoute(route, view, action string){
	http.HandleFunc(route,router)
	view_name = view
	action_name = action
}

func router(w http.ResponseWriter, r *http.Request){
	controller_ref := controller.Controller{}
	params := r.URL.Query(); //get all params in URL
	fmt.Println("Parameters received: ", params)
	t, err := template.ParseFiles(view_name)
	if err != nil {
        panic(err)
    }	
    in := []reflect.Value{reflect.ValueOf(params)}
    return_values := reflect.ValueOf(&controller_ref).MethodByName(action_name).Call(in) //pass params to controller method and return view_vars
    i := return_values[0].Interface()
    view_vars := i.(map[string]string)
	t.Execute(w, view_vars)
	fmt.Println("--> Served" , r.URL)
}

func Init(){
	fmt.Println("--> Listening on port 9000")
	http.ListenAndServe(":9000",nil)
}
