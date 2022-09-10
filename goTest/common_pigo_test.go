package goTest

import (
	pigo "github.com/esimov/pigo/core"
	"github.com/zeromicro/go-zero/core/logx"
	"gocv.io/x/gocv"
	"io/ioutil"
	"testing"
)

func TestPigo(t *testing.T) {
	cascadeFile, err := ioutil.ReadFile("C:\\Users\\戴达\\go\\pkg\\mod\\github.com\\esimov\\pigo@v1.4.5\\cascade\\facefinder")
	if err != nil {
		logx.Errorf("Error reading the cascade file: %v", err)
	}

	src, err := pigo.GetImage("C:\\Users\\戴达\\go\\pkg\\mod\\github.com\\esimov\\pigo@v1.4.5\\testdata\\sample.jpg")
	if err != nil {
		logx.Errorf("Cannot open the image file: %v", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,

		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	pigo := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pigo.Unpack(cascadeFile)
	if err != nil {
		logx.Errorf("Error reading the cascade file: %s", err)
	}

	angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians

	// Run the classifier over the obtained leaf nodes and return the detection results.
	// The result contains quadruplets representing the row, column, scale and detection score.
	dets := classifier.RunCascade(cParams, angle)

	logx.Info(dets)

	// Calculate the intersection over union (IoU) of two clusters.
	dets = classifier.ClusterDetections(dets, 0.2)

	logx.Info(dets)

	//plc, err = pl.UnpackCascade(flpc)
	//if err != nil {
	//	t.Fatalf("failed unpacking the cascade file: %v", err)
	//}
	//
	//for _, det := range dets {
	//	if det.Scale > 50 {
	//		// left eye
	//		puploc := &pigo.Puploc{
	//			Row:      det.Row - int(0.075*float32(det.Scale)),
	//			Col:      det.Col - int(0.175*float32(det.Scale)),
	//			Scale:    float32(det.Scale) * 0.25,
	//			Perturbs: 50,
	//		}
	//		leftEye := plc.RunDetector(*puploc, *imgParams, 0.0, false)
	//
	//		// right eye
	//		puploc = &pigo.Puploc{
	//			Row:      det.Row - int(0.075*float32(det.Scale)),
	//			Col:      det.Col + int(0.185*float32(det.Scale)),
	//			Scale:    float32(det.Scale) * 0.25,
	//			Perturbs: 50,
	//		}
	//		rightEye := plc.RunDetector(*puploc, *imgParams, 0.0, false)
	//
	//		flp := plc.GetLandmarkPoint(leftEye, rightEye, *imgParams, perturb, false)
	//		landMarkPoints = append(landMarkPoints, *flp)
	//	}
	//}
	//if len(landMarkPoints) == 0 {
	//	t.Fatal("should have been detected facial landmark points")
	//}
}

func TestPigoLocalhost(t *testing.T) {
	//webcam, err := gocv.VideoCaptureDevice(0)
	//
	//if nil != err {
	//	fmt.Println("VideoCaptureDevice err ", err)
	//	return
	//}
	//defer webcam.Close()
	//
	//window := gocv.NewWindow("pigo")
	//defer window.Close()
	//
	//window.ResizeWindow(640, 480)
	//
	//img := gocv.NewMat()
	//defer img.Close()
	//
	//green := color.RGBA{0, 255, 0, 0}
	//
	//cascadeFile, err := ioutil.ReadFile("facefinder")
	//
	//if err != nil {
	//	logx.Errorf("Error reading the cascade file: %v", err)
	//}
	//
	//for {
	//
	//	if ok := webcam.Read(&img); !ok {
	//		fmt.Println("Read err ")
	//		return
	//	}
	//
	//	if img.Empty() {
	//		continue
	//	}
	//
	//	goImg, err := img.ToImage()
	//
	//	if nil != err {
	//		logx.Errorf("ToImage err ")
	//		return
	//	}
	//
	//	pixels := pigo.RgbToGrayscale(goImg)
	//	cols, rows := goImg.Bounds().Max.X, goImg.Bounds().Max.Y
	//
	//	cParams := pigo.CascadeParams{
	//		MinSize:     20,
	//		MaxSize:     1000,
	//		ShiftFactor: 0.1,
	//		ScaleFactor: 1.1,
	//
	//		ImageParams: pigo.ImageParams{
	//			Pixels: pixels,
	//			Rows:   rows,
	//			Cols:   cols,
	//			Dim:    cols,
	//		},
	//	}
	//
	//	pPigo := pigo.NewPigo()
	//
	//	classifier, err := pPigo.Unpack(cascadeFile)
	//	if err != nil {
	//		logx.Errorf("Error reading the cascade file: %s", err)
	//	}
	//
	//	angle := 0.0
	//	iouThreshold := 0.3
	//
	//	dets := classifier.RunCascade(cParams, angle)
	//
	//	dets = classifier.ClusterDetections(dets, iouThreshold)
	//
	//	for _, face := range dets {
	//
	//		if face.Q > 5 {
	//			x := face.Col - face.Scale/2
	//			y := face.Row - face.Scale/2
	//			r := image.Rect(x, y, x+face.Scale, y+face.Scale)
	//			gocv.Rectangle(&img, r, green, 3)
	//
	//		} else {
	//			continue
	//		}
	//
	//	}
	//
	//	window.IMShow(img)
	//
	//	if 27 == window.WaitKey(1) {
	//		break
	//	}
	//
	//}
	gocv.NewWindow("Hello")
}
