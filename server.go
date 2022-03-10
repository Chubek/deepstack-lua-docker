package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
)

func decodeBase64(text string) string {
	dec, _ := base64.StdEncoding.DecodeString(text)

	return string(dec)

}

func hello(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	command := queries.Get("command")
	args := queries.Get("args")

	fmt.Println(command, args)

	out, err := exec.Command(command, args).Output()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	w.Write([]byte(fmt.Sprintf("%s\n%s", "Executed Successfully, Results: ", out)))
}

func html(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(decodeBase64(`PCFET0NUWVBFIGh0bWw+PGh0bWw+PGhlYWQ+IDx0aXRsZT5UaGluZ2FtYWplZWs8L3RpdGxlPiA8bGluayByZWw9InN0eWxlc2hlZXQiIGhyZWY9Imh0dHBzOi8vY2RuanMuY2xvdWRmbGFyZS5jb20vYWpheC9saWJzL2RhaXN5dWkvMS4yNS40L2Z1bGwubWluLmNzcyIgaW50ZWdyaXR5PSJzaGE1MTItZkxkUFRzZHRBeE4rL3llektNWUZKV1RHVUVpRW5JRkhVN1BQdDBhL1ZoTE9VV1BoRnNXYVdzSTJjSmdGUS9rcFR3NUFwTUo4QzRadytiZXVRZHZGMlE9PSIgY3Jvc3NvcmlnaW49ImFub255bW91cyIgcmVmZXJyZXJwb2xpY3k9Im5vLXJlZmVycmVyIi8+IDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cHM6Ly9jZG5qcy5jbG91ZGZsYXJlLmNvbS9hamF4L2xpYnMvdGFpbHdpbmRjc3MvMi4yLjE5L3RhaWx3aW5kLm1pbi5jc3MiIGludGVncml0eT0ic2hhNTEyLXduZWE5OXVLSUMzVEpGN3Y0ZUtrNFkrbE16Mk1rbHYxOCtyNG5hMkduMWFiRFJQUE9lZWY5NXhUemR3R0Q5ZTZ6WEpCdGVNSWhaMSs2OFFDNWJ5Slp3PT0iIGNyb3Nzb3JpZ2luPSJhbm9ueW1vdXMiIHJlZmVycmVycG9saWN5PSJuby1yZWZlcnJlciIvPiA8c3R5bGU+LmhvdmVyLWNsb3NlOmhvdmVye2NvbG9yOiByZ2IoMjU1LCAyNTUsIDI1NSk7IGN1cnNvcjogcG9pbnRlcjt9PC9zdHlsZT48L2hlYWQ+PGJvZHk+IDxkaXYgaWQ9ImFwcCIgY2xhc3M9ImdyaWQgdy1zY3JlZW4gaC1zY3JlZW4gYmctYmFzZS0zMDAgcGxhY2UtaXRlbXMtY2VudGVyIHJvdW5kZWQtbWQgaG92ZXI6cm91bmRlZC1sZyI+IDxkaXYgY2xhc3M9ImZsZXggZmxleC1jb2wgdy0zLzQgaC01MCBnYXAtMyI+IDxkaXYgY2xhc3M9ImZsZXggZmxleC1yb3ciPiA8ZGl2IGNsYXNzPSJpbnB1dC1ncm91cCB3LTIvNCI+IDxzcGFuPkNvbW1hbmQ8L3NwYW4+IDx0ZXh0YXJlYSB2LW1vZGVsPSJjb21tYW5kIiBjbGFzcz0idGV4dGFyZWEgdGV4dGFyZWEtZXJyb3Igdy1mdWxsIiBwbGFjZWhvbGRlcj0iQ29tbWFuZCI+PC90ZXh0YXJlYT4gPC9kaXY+PGRpdiBjbGFzcz0iaW5wdXQtZ3JvdXAgdy0yLzQiPiA8c3Bhbj5BcmdzPC9zcGFuPiA8dGV4dGFyZWEgdi1tb2RlbD0iYXJncyIgY2xhc3M9InRleHRhcmVhIHRleHRhcmVhLWVycm9yIHctZnVsbCIgcGxhY2Vob2xkZXI9IkFyZ3MiPjwvdGV4dGFyZWE+IDwvZGl2PjwvZGl2PjxidXR0b24gQGNsaWNrPSJydW4iIGNsYXNzPSJidG4gYnRuLWxnIj5SdW48L2J1dHRvbj4gPGRpdiBjbGFzcz0ibW9ja3VwLWNvZGUiPiA8cHJlPjxjb2RlPnt7IHJlc3VsdCB9fTwvY29kZT48L3ByZT4gPC9kaXY+PC9kaXY+PC9kaXY+PHNjcmlwdCBzcmM9Imh0dHBzOi8vY2RuLmpzZGVsaXZyLm5ldC9ucG0vdnVlQDIuNi4xNC9kaXN0L3Z1ZS5qcyI+PC9zY3JpcHQ+IDxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0Ij5sZXQgdnVlPW5ldyBWdWUoe2VsOiAiI2FwcCIsIGRhdGEoKXtyZXR1cm57YXJnczogIm1hdGNoTmFtZSBob2xkZW0ubGltaXQuMnAucmV2ZXJzZV9ibGluZHMuZ2FtZSAxMDAwIDAgQWxpY2UgLi9leGFtcGxlX3BsYXllci5saW1pdC4ycC5zaCBCb2IgLi9leGFtcGxlX3BsYXllci5saW1pdC4ycC5zaCIsIHJlc3VsdDogIlJlc3VsdCB3aWxsIHNob3cgdXAgaGVyZS4uLiIsIGNvbW1hbmQ6ICJjZCB+L0RlZXBIb2xkZW0vQUNQQ1NlcnZlciAmJiAuL3BsYXlfbWF0Y2gucGwifX0sIGNyZWF0ZWQ6IGZ1bmN0aW9uKCl7fSwgbW91bnRlZCgpe30sIGNvbXB1dGVkOnt9LCB3YXRjaDp7fSwgbWV0aG9kczp7cnVuOiBhc3luYyAoKT0+e2NvbnN0IGNvbW1FbmM9ZW5jb2RlVVJJQ29tcG9uZW50KHZ1ZS5jb21tYW5kKTsgY29uc3QgYXJnRW5jPWVuY29kZVVSSUNvbXBvbmVudCh2dWUuYXJncyk7IGNvbnN0IHJldFRleHQ9YXdhaXQgKGF3YWl0IGZldGNoKCBgL3J1bj9jb21tYW5kPSR7Y29tbUVuY30mYXJncz0ke2FyZ0VuY31gICkpLnRleHQoKTsgdnVlLnJlc3VsdD1yZXRUZXh0O319fSkgPC9zY3JpcHQ+PC9ib2R5PjwvaHRtbD4=`)))
}

func main() {

	http.HandleFunc("/run", hello)
	http.HandleFunc("/", html)
	fmt.Println("Serving... on port 8090")
	http.ListenAndServe(":8090", nil)
}