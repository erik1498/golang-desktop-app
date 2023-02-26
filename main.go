package main

import "github.com/webview/webview"

func main() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Sekolah App")
	w.SetSize(800, 600, webview.Hint(webview.HintNone))
	w.Navigate("http://localhost:4000/kurikulum/swagger-ui/#/")
	w.Run()
}
