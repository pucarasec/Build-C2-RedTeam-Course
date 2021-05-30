package layer0

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"../handler"
)

type HTTPHandler struct {
	subhandler handler.Handler
}

func NewHTTPHandler(subhandler handler.Handler) *HTTPHandler {
	return &HTTPHandler{subhandler: subhandler}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodedIncomingMsg := r.PostFormValue("m")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(encodedIncomingMsg) == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	incomingMsg, err := base64.StdEncoding.DecodeString(encodedIncomingMsg)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response, err := h.subhandler.HandleMsg(incomingMsg)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(response)

	fmt.Fprintf(w, "<!DOCTYPE html><html lang=\"en\">")
	fmt.Fprintf(w, "<body>Hi there, this is a totally innocent web page. <!--%s--></body>", encoded)
	fmt.Fprintf(w, "</html>")
}
