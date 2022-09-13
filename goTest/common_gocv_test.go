package goTest

import (
	"errors"
	"fmt"
	"github.com/daida459031925/common/error/try"
	"github.com/daida459031925/common/reflex"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"reflect"
	"strconv"
	"strings"
	"strings"
	"testing"
)

type (
	face struct {
		imgFace     gocv.Mat                 //需要检测的人脸原始照片 裁剪
		imgCropFace gocv.Mat                 //裁剪成人脸的照片
		classifier  []gocv.CascadeClassifier //加载分类器xml对象
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

// 初始化face struct
func NewFace() face {
	return face{closeStruct: closeStruct{glist.New()}}
}

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

// 加载分类器，目前包括人脸，人眼
// haarcascade_frontalface_alt.xml haarcascade_eye_tree_eyeglasses.xml
func (f face) loadFaceXmlFile(xml ...string) {
	fx.From(func(source chan<- interface{}) {
		if xml != nil && len(xml) > 0 {
			for i := range xml {
				source <- strings.TrimSpace(xml[i])
			}
		}
	}).Distinct(func(item interface{}) interface{} {
		s := item.(string)
		return s
	}).ForEach(func(item interface{}) {
		s := item.(string)
		logx.Info(s)
		if len(s) > 0 {
			classifier := gocv.NewCascadeClassifier()
			if !classifier.Load(s) {
				logx.Errorf("load %s fail", s)
				_ = classifier.Close()
				return
			}
			f.classifier = append(f.classifier, classifier)
			f.closeList.PushBack(classifier.Close())
		}
	})
}

func a(a gocv.CascadeClassifier) {

}

func TestInit(t *testing.T) {
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

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	for _, face := range imageFaces {
		gocv.Rectangle(&img, face, blue, 3)

		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(face.Min.X+(face.Min.X/2)-(size.X/2), face.Min.Y-2)
		gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	}

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
	//imageFace := "C:\\Users\\daida\\Desktop\\14EA0B1FBEB616B217BBCA8B85462C9E.jpg"
	imageFace := "C:\\Users\\daida\\Desktop\\WIN_20220903_18_50_18_Pro.jpg"
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
	window.ResizeWindow(face.Max.X-face.Min.X, face.Max.Y-face.Min.Y)
	window.IMShow(croppedMat)
	//防止程序关闭 按esc关闭
	for {
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

// 笔记本摄像头人脸检测
func TestFaceCamera(t *testing.T) {
	faceXml := "D:\\opencv\\sources\\data\\haarcascades\\haarcascade_frontalface_alt.xml"
	eyeXml := "D:\\opencv\\sources\\data\\haarcascades\\haarcascade_eye_tree_eyeglasses.xml"
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	//load classifier to recognize faces
	if !classifier.Load(faceXml) {
		logx.Error("faceXml load err")
		exit()
	}

	eyeclassifier := gocv.NewCascadeClassifier()
	defer eyeclassifier.Close()
	//load classifier to recognize faces
	if !eyeclassifier.Load(eyeXml) {
		logx.Error("eyeXml load err")
		exit()
	}

	webcam, _ := gocv.VideoCaptureDevice(0)
	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}
	img := gocv.NewMat()
	// show the image in the window, and wait 1 millisecond
	for {
		if ok := webcam.Read(&img); !ok {
			return
		}
		if img.Empty() {
			continue
		}
		imageFaces := classifier.DetectMultiScale(img)
		faces := len(imageFaces)
		if faces > 0 {
			logx.Info("检测到人脸")
			for _, face := range imageFaces {
				gocv.Rectangle(&img, face, blue, 3)

				size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
				pt := image.Pt(face.Min.X+(face.Min.X/2)-(size.X/2), face.Min.Y-2)
				gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
			}
		}

		eyes := eyeclassifier.DetectMultiScale(img)
		eye := len(eyes)
		if eye > 0 {
			logx.Info("检测到眼睛")
			for _, face := range eyes {
				gocv.Rectangle(&img, face, blue, 3)

				size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
				pt := image.Pt(face.Min.X+(face.Min.X/2)-(size.X/2), face.Min.Y-2)
				gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
			}
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}

func TestFaceStruct(t *testing.T) {
	f := NewFace()
	f.loadFaceXmlFile("D:\\opencv\\sources\\data\\haarcascades\\haarcascade_frontalface_alt.xml",
		"D:\\opencv\\sources\\data\\haarcascades\\haarcascade_eye_tree_eyeglasses.xml",
		"D:\\opencv\\sources\\data\\haarcascades\\haarcascade_frontalface_alt.xml",
		"D:\\opencv\\sources\\data\\haarcascades\\haarcascade_eye_tree_eyeglasses.xml")
}

func TestMath(t *testing.T) {
	a, e := strconv.ParseFloat("", 64)
	if e != nil {
		logx.Info(a)
	}

	fff := -8.0497183772403904e-17

	logx.Info(decimal.NewFromFloat(fff))

	logx.Info(fmt.Sprintf("%f", fff))

	d, _ := getDecimal("-2.22222222222222222222222222222222222222")
	b, _ := getDecimal("-0.00000000000000000000000000000000000001")
	b0, _ := getDecimal("2")
	d = d.Add(b)

	logx.Info(d.Add(d).DivRound(b0, 40))
}

func getDecimal(str string) (decimal.Decimal, error) {
	s := strings.TrimSpace(str)
	_, e := strconv.ParseFloat(str, 64)
	if e != nil {
		s = "0"
		e = errors.New(fmt.Sprintf("转换float64失败: %s", e))
	}

	return decimal.RequireFromString(s), e
}
