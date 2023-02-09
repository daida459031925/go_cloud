package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
	"time"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
	//test("C:\\Users\\daida\\Desktop\\cn_sinacn10_309_w534h575_20180926_a837-hhuhisn1021919.jpg", "C:\\Users\\daida\\Desktop\\6HX3N4ZD2LF`77)JW5)$~LD.png")

	count := 0

	for i := 1; i < 1026; i++ {
		for i1 := 0; i1 < i; i1++ {
			for j := 0; j < 256^i; j++ {
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println(math.Pow(256, 3))
}

// 图片查找算法，在大图片中查找小图片；若存在相似度高的则返还则展示否则退出。
// 单独一张静态图片和一张需要找的静态图片，毫秒数为100以内
func test(srcPath, templatePath string) {
	l := time.Now().UnixMilli()

	src := srcPath           //匹配图片
	template := templatePath //模板图片

	imgTempl := gocv.IMRead(template, gocv.IMReadGrayScale)
	if imgTempl.Empty() {
		logx.Info("Invalid read of %s", template)
		return //为空，返回
	}
	defer imgTempl.Close()

	imgSrc := gocv.IMRead(src, gocv.IMReadGrayScale)
	if imgSrc.Empty() {
		logx.Info("Invalid read of %s", src)
		return //为空，返回
	}
	defer imgSrc.Close()

	imgSrc1 := gocv.IMRead(src, gocv.IMReadColor)
	if imgSrc1.Empty() {
		logx.Info("Invalid read of %s", src)
		return //为空，返回
	}
	defer imgSrc1.Close()

	window := gocv.NewWindow("Image Match")
	defer window.Close()

	result := gocv.NewMat()
	defer result.Close()

	m := gocv.NewMat()

	blue := color.RGBA{0, 0, 255, 0}
	gocv.MatchTemplate(imgTempl, imgSrc, &result, gocv.TmCcoeffNormed, m)
	m.Close()
	_, maxConfidence, _, maxLoc := gocv.MinMaxLoc(result)
	if maxConfidence < 0.95 {
		logx.Infof("Max confidence of %f is too low. MatchTemplate could not find template in scene.", maxConfidence)
		return
	}
	rect := image.Rect(maxLoc.X, maxLoc.Y, maxLoc.X+imgTempl.Cols(), maxLoc.Y+imgTempl.Rows())
	gocv.Rectangle(&imgSrc1, rect, blue, 3)
	l1 := time.Now().UnixMilli()
	logx.Info(l1)
	logx.Info(l)
	logx.Infof("l1-l:= %d", l1-l)
	//for {
	//	window.IMShow(imgSrc1)
	//	if window.WaitKey('q') >= 0 {
	//		break
	//	}
	//}

}

// 旋转90度
func rotate90(m image.Image) image.Image {
	rotate90 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			//  设置像素点
			rotate90.Set(m.Bounds().Max.Y-x, y, m.At(y, x))
		}
	}
	return rotate90
}

// 旋转180度

func rotate180(m image.Image) image.Image {
	rotate180 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dx(), m.Bounds().Dy()))
	// 矩阵旋转
	for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
		for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
			//  设置像素点
			rotate180.Set(m.Bounds().Max.X-x, m.Bounds().Max.Y-y, m.At(x, y))
		}
	}
	return rotate180
}

// 旋转270度
func rotate270(m image.Image) image.Image {
	rotate270 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			// 设置像素点
			rotate270.Set(x, m.Bounds().Max.X-y, m.At(y, x))
		}
	}
	return rotate270

}

// 将图片居中处理
func centerImage(m image.Image) image.Image {
	// 现在图片是长>宽，将图片居中设置
	max := m.Bounds().Dx()
	// 居中后距离最底部的高度为(x-y)/2
	temp := (max - m.Bounds().Dy()) / 2
	centerImage := image.NewRGBA(image.Rect(0, 0, max, max))
	for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
		for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
			centerImage.Set(x, temp+y, m.At(x, y))
		}
	}
	return centerImage

}
