package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	video "app/shared/video/ffmpeg"
	"app/shared/view"
)

type Frame struct {
	imageData string
}

// AboutGET displays the About page
func AboutGET(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "about/about"
	v.Render(w)
}

// VideotoJPEG converts the VIDEO to JPEG with the frame number
func VideotoJPEG(w http.ResponseWriter, r *http.Request) {

	fmt.Println("PATH")
	urlParams := strings.Split(r.URL.Path, "&!")
	fmt.Println(urlParams)
	assetURL := "blob:http://localhost:8080/" + urlParams[1]
	framenumber, _ := strconv.Atoi(urlParams[2])
	buf := video.ExampleReadFrameAsJpeg(assetURL, framenumber)
	message := fmt.Sprintf(" { \"Messages\" : \"%s\" }", buf)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, message)

}

// Encodex264 converts the VIDEO to JPEG with the frame number
func Encodex264(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 10)

	// decoder := json.NewDecoder(r.Body)
	// var t Frame
	// err := decoder.Decode(&t)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(t.imageData)

	fmt.Println("*************************************************************************************")
	frmCategory := r.PostFormValue("category")
	fmt.Println(frmCategory)
	fmt.Println("*************************************************************************************")
	// // Declare a new Person struct.
	// var f Frame
	// time.Sleep(time.Second * 10)
	// // Try to decode the request body into the struct. If there is an error,
	// // respond to the client with the error message and a 400 status code.
	// err := json.NewDecoder(r.Body).Decode(&f)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // Do something with the Person struct...
	// fmt.Println("*************************************************************************************")
	// fmt.Println("ImageDataStart")
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(r.Body)

	// //Convert the body to type string
	// sb := string(body)
	// log.Printf(sb)

	// fmt.Println("ImageDataStart")
	// fmt.Println("*************************************************************************************")
	// opts := &x264.Options{
	// 	Width:     640,
	// 	Height:    480,
	// 	FrameRate: 25,
	// 	Preset:    "veryfast",
	// 	Tune:      "stillimage",
	// 	Profile:   "baseline",
	// 	LogLevel:  x264.LogDebug,
	// }

	// // w, err := os.Create("example.264")
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	os.Exit(1)
	// // }

	// // defer w.Close()

	// enc, err := x264.NewEncoder(w, opts)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// defer enc.Close()

	// // r, err = os.Open("example.jpg")
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	os.Exit(1)
	// // }

	// img, _, err := image.Decode(r)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// err = enc.Encode(img)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// err = enc.Flush()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println("PATH")
	// urlParams := strings.Split(r.URL.Path, "&!")
	// fmt.Println(urlParams)
	// assetURL := "blob:http://localhost:8080/" + urlParams[1]
	// framenumber, _ := strconv.Atoi(urlParams[2])
	// buf := video.ExampleReadFrameAsJpeg(assetURL, framenumber)
	// message := fmt.Sprintf(" { \"Messages\" : \"%s\" }", buf)
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, frmCategory)

}
