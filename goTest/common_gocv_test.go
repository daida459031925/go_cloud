package goTest

import (
	"github.com/daida459031925/common/error/try"
	"github.com/daida459031925/common/reflex"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
	"reflect"
	"testing"
)

type (
	face struct {
		imgFace     gocv.Mat               //需要检测的人脸原始照片 裁剪
		imgCropFace gocv.Mat               //裁剪成人脸的照片
		classifier  gocv.CascadeClassifier //加载人脸分类器xml对象
		closeStruct
	}

	closeStruct struct {
		closeList *glist.List //需要关闭程序的集合对象，目标是关闭操作对象集合，由最后一个添加进入的方法进行关闭
	}

	//为了防止编写两个很重复的方法使用组合对象实现防止重复代码过多
	pop interface {
		getFunc(l *glist.List) interface{}
	}

	popFront struct {
	}

	popBack struct {
	}
)

// 为了减少重复代码编写以及维护难度提高方法只写一处，执行方法使用pop组合对象方式减少代码重复度
func execute(g *glist.List, p pop) {
	v := reflex.GetRef(p).GetPointerData(false).Type()
	fx.From(func(source chan<- interface{}) {
		if g != nil && g.Len() > 0 {
			for i := 0; i < g.Len(); i++ {
				source <- p.getFunc(g)
			}
		}
	}).ForEach(func(item interface{}) {
		r, e := item.(func())
		if e {
			try.Try(r).CatchAll(func(err error) {
				logx.Errorf("to by %v type , server func error: %s", v, err.Error())
			})
		} else {
			logx.Errorf("to by %v type , server is not func error", v)
		}
	})
}

// 以栈的模式先进后出
func (cs closeStruct) executePopBack() {
	execute(cs.closeList, new(popBack))
}

// 以队列的模式先进先出
func (cs closeStruct) executePopFront() {
	execute(cs.closeList, new(popFront))
}

func (popFront) getFunc(l *glist.List) interface{} {
	return l.PopFront()
}

func (popBack) getFunc(l *glist.List) interface{} {
	return l.PopBack()
}

func loadFile(faceXml string, imageFace string) {

}

func TestInit(t *testing.T) {
	//webcam, _ := gocv.VideoCaptureDevice(0)
	//window := gocv.NewWindow("Hello")
	//img := gocv.NewMat()
	//for {
	//	webcam.Read(&img)
	//	window.IMShow(img)
	//	window.WaitKey(1)
	//}
	faceXml := "D:\\opencv\\sources\\data\\haarcascades\\haarcascade_frontalface_alt.xml"
	imageFace := "C:\\Users\\daida\\Desktop\\src=http___n.sinaimg.cn_sinacn10_309_w534h575_20180926_a837-hhuhisn1021919.jpg&refer=http___n.sinaimg.jpg"
	//imageFace := "C:\\Users\\daida\\go\\pkg\\mod\\gocv.io\\x\\gocv@v0.31.0\\images\\aruco_6X6_250_6.png"
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	//load classifier to recognize faces
	if !classifier.Load(faceXml) {
		logx.Error("faceXml load err")
		exit()
	}
	mat := gocv.NewMat()
	defer mat.Close()
	img := gocv.IMRead(imageFace, gocv.IMReadUnchanged)
	imageFaces := classifier.DetectMultiScale(img)
	faces := len(imageFaces)
	if faces <= 0 {
		logx.Error("未检测到人脸")
		exit()
	}
	logx.Infof("检测到 %d 个人脸", faces)

	// open display window
	//window := gocv.NewWindow("Face Detect")
	//defer window.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	for _, face := range imageFaces {
		gocv.Rectangle(&img, face, blue, 3)

		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(face.Min.X+(face.Min.X/2)-(size.X/2), face.Min.Y-2)
		gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	}

	// show the image in the window, and wait 1 millisecond
	//window.IMShow(img)
	//for {
	//	if window.WaitKey(1) >= 0 {
	//		break
	//	}
	//}

}

// 使用gocv中所有可以获取检测图像关键点的方法
func TestFaceDetectAll(t *testing.T) {
	logx.Infof("%v", reflect.ValueOf(new(popBack)).Elem().Type())
	logx.Infof("%v", reflex.GetRef(new(popBack)).GetPointerData(false).Type())
	logx.Infof("%v", reflex.GetRef(new(popBack)).GetPointerData(true).Type())

	imageFace := "C:\\Users\\daida\\Desktop\\src=http___n.sinaimg.cn_sinacn10_309_w534h575_20180926_a837-hhuhisn1021919.jpg&refer=http___n.sinaimg.jpg"
	img := gocv.IMRead(imageFace, gocv.IMReadUnchanged)

	akaze := gocv.NewAKAZE()
	defer akaze.Close()
	akazes := akaze.Detect(img)
	//for _, point := range akazes {
	//	logx.Info(point)
	//}
	logx.Info(len(akazes))

	afd := gocv.NewAgastFeatureDetector()
	defer afd.Close()
	afds := afd.Detect(img)
	//for _, point := range afds {
	//	logx.Info(point)
	//}
	logx.Info(len(afds))

	brisk := gocv.NewBRISK()
	defer brisk.Close()
	brisks := brisk.Detect(img)
	logx.Info(len(brisks))

	gfttd := gocv.NewGFTTDetector()
	defer gfttd.Close()
	gfttds := gfttd.Detect(img)
	logx.Info(len(gfttds))

	kaze := gocv.NewKAZE()
	defer kaze.Close()
	kazes := kaze.Detect(img)
	logx.Info(len(kazes))

	ffd := gocv.NewFastFeatureDetector()
	defer ffd.Close()
	ffds := ffd.Detect(img)
	logx.Info(len(ffds))

	mser := gocv.NewMSER()
	defer mser.Close()
	msers := mser.Detect(img)
	logx.Info(len(msers))

	ORB := gocv.NewORB()
	defer ORB.Close()
	ORBs := ORB.Detect(img)
	logx.Info(len(ORBs))

	sbd := gocv.NewSimpleBlobDetector()
	defer sbd.Close()
	sbds := sbd.Detect(img)
	logx.Info(len(sbds))

	SIFT := gocv.NewSIFT()
	defer SIFT.Close()
	SIFTs := SIFT.Detect(img)
	logx.Info(len(SIFTs))

}

// faceCompute
func TestFaceCompute(t *testing.T) {
	imageFaceSrc := "C:\\Users\\daida\\Desktop\\src=http___n.sinaimg.cn_sinacn10_309_w534h575_20180926_a837-hhuhisn1021919.jpg&refer=http___n.sinaimg.jpg"
	imageFaceCompute := "C:\\Users\\daida\\Desktop\\14EA0B1FBEB616B217BBCA8B85462C9E.jpg"
	//imageFaceCompute := "C:\\Users\\daida\\Desktop\\src=http___n.sinaimg.cn_sinacn10_309_w534h575_20180926_a837-hhuhisn1021919.jpg&refer=http___n.sinaimg.jpg"

	imgSrc := gocv.IMRead(imageFaceSrc, gocv.IMReadUnchanged)
	imgCompute := gocv.IMRead(imageFaceCompute, gocv.IMReadUnchanged)

	bf := gocv.NewBFMatcher()
	defer bf.Close()

	ORB := gocv.NewORB()
	defer ORB.Close()

	i0 := gocv.NewMat()
	defer i0.Close()

	i1 := gocv.NewMat()
	defer i1.Close()

	_, is := ORB.DetectAndCompute(imgSrc, i0)
	_, ic := ORB.DetectAndCompute(imgCompute, i1)
	//knn筛选结果
	knn := bf.KnnMatch(is, ic, 2)
	good := make([][]gocv.DMatch, 0)
	for _, matches := range knn {
		if matches[0].Distance < 0.85*matches[1].Distance {
			good = append(good, matches)
		}
	}
	similary := float64(len(good)) / float64(len(knn))

	logx.Infof("(ORB算法)两张图片相似度为: %f %s", similary*100, "%")

}

// faceCutOut
func TestFaceCutOut(t *testing.T) {
	//加载人脸配置文件
	faceXml := "D:\\opencv\\sources\\data\\haarcascades\\haarcascade_frontalface_alt.xml"
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	classifier.Load(faceXml)

	//加载需要读取的图片资源
	imageFace := "C:\\Users\\daida\\Desktop\\14EA0B1FBEB616B217BBCA8B85462C9E.jpg"
	img := gocv.IMRead(imageFace, gocv.IMReadUnchanged)

	//检查是否有人脸
	imageFaces := classifier.DetectMultiScale(img)
	faces := len(imageFaces)
	if faces <= 0 {
		logx.Error("未检测到人脸")
		exit()
	}
	logx.Infof("检测到 %d 个人脸", faces)

	face := imageFaces[0]

	croppedMat := img.Region(image.Rect(face.Min.X, face.Min.Y, face.Max.X, face.Max.Y))

	//打开windows 查看图片
	window := gocv.NewWindow("Face Detect")
	defer window.Close()
	//设置大小并计算画面大小二项坐标 坐标点A（左上）,B（右下）；X的差值 Y的差值
	window.ResizeWindow(math.Abs(face.Max.X-face.Min.X), face.Max.Y-face.Min.Y)
	window.IMShow(croppedMat)
	//防止程序关闭 按esc关闭
	for {
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
