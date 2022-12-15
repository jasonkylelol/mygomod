package mygomod

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func buildImgData(imgName, imgsPath string) (string, error) {
	imgFile := fmt.Sprintf("%v/%v", imgsPath, imgName)
	imgData, err := ioutil.ReadFile(imgFile)
	if err != nil {
		fmt.Println("ioutil.ReadFile failed", err)
		return "", err
	}
	// var base64Encoding strings.Builder
	// // Determine the content type of the image file
	// mimeType := http.DetectContentType(imgData)
	// // Prepend the appropriate URI scheme header depending
	// // on the MIME type
	// switch mimeType {
	// case "image/jpeg":
	// 	fmt.Fprintf(&base64Encoding, "%v", "data:image/jpeg;base64,")
	// case "image/png":
	// 	fmt.Fprintf(&base64Encoding, "%v", "data:image/png;base64,")
	// }
	// // Append the base64 encoded output
	// fmt.Fprintf(&base64Encoding, "%v", base64.StdEncoding.EncodeToString(imgData))
	// base64Encoding += base64.StdEncoding.EncodeToString(imgData)
	return base64.StdEncoding.EncodeToString(imgData), nil
}
