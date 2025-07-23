package ipfs

import (
	"fmt"

	"github.com/zde37/pinata-go-sdk/pinata"
)


func UploadToIPFS() {

	// Expose the jwt cuz why not :)
	auth := pinata.NewAuthWithJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiJkOTc3YTllYy04MGIxLTQwY2MtOWY0Ny0zNDllNTA0NDFhMTMiLCJlbWFpbCI6ImpkZG9raWppaWZmdUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicGluX3BvbGljeSI6eyJyZWdpb25zIjpbeyJkZXNpcmVkUmVwbGljYXRpb25Db3VudCI6MSwiaWQiOiJGUkExIn0seyJkZXNpcmVkUmVwbGljYXRpb25Db3VudCI6MSwiaWQiOiJOWUMxIn1dLCJ2ZXJzaW9uIjoxfSwibWZhX2VuYWJsZWQiOmZhbHNlLCJzdGF0dXMiOiJBQ1RJVkUifSwiYXV0aGVudGljYXRpb25UeXBlIjoic2NvcGVkS2V5Iiwic2NvcGVkS2V5S2V5IjoiZjVjNDBiY2ViNzVlMWZjMjExNGMiLCJzY29wZWRLZXlTZWNyZXQiOiIxZGNlYzMzYjNmZGMzZjc1OGM5ZTM5NGM4YWYzNzgwMWU2ZGZmN2E0MjY3M2MwNjNhOWMwOWQ2ZWNmZWE3ZmZjIiwiZXhwIjoxNzg0ODI4NzQ0fQ.kFVAcu0HqdECdgo_NV8jhPzPwlOnS0bvDL8WM1BL8zY")
	client := pinata.New(auth)

	response, err := client.PinFile("/home/linuxer77/t.json", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File pinned successfully. IPFS hash: %s\n", response)

}