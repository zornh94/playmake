package transport
import (
"fmt"
"net/http"
)
func StartServer(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w,"%s","hello")
})
http.ListenAndServe(":8090", nil)
}
