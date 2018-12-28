# FaceDetection-Using-Open-CV-



## To run Before
```
source secure-key

docker run -d -p 8080:8080 -e "MB_KEY=$MB_KEY" machinebox/facebox

```

## To run
```
go run main.go
```

## Video Capture

```
webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
  
  ```
  
  ### Reference 
  * [Open CV doc](https://opencv.org)
  * [1](https://gocv.io)
  * [2](https://github.com/go-opencv/go-opencv)
