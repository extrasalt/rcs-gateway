package main 

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main(){

    r:= mux.NewRouter()
    r.HandleFunc("/file", FileUploadHandler).Methods("POST")
    

    http.ListenAndServe(":3000", r)
}


func FileUploadHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("File upload handler"))
}
