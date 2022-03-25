/*
@Description:
@File : pdf
@Author : blue_sky_12138
@Version: 1.0.0
@Date : 2021/12/23 13:50
*/

package pdf

import (
	"flag"
	"fmt"
	"github.com/signintech/gopdf"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Size struct {
	Height float64
	Weight float64
}

func PdfImages(dicPath string, size ...Size) {
	pdf := gopdf.GoPdf{}
	si := *gopdf.PageSizeA4
	if len(size) != 0 {
		si.W = size[0].Weight
		si.H = size[0].Height
	}
	pdf.Start(gopdf.Config{PageSize: si})

	dir, err := ioutil.ReadDir(dicPath)
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		if file.IsDir() || file.Name()[len(file.Name())-3:] == ".db" {
			continue
		}

		path := dicPath + string(os.PathSeparator) + file.Name()

		pdf.AddPage()
		err = pdf.Image(path, 0, 0, nil)

		if err != nil {
			log.Println(path)
			panic(err)
		}
	}

	splits := strings.Split(dicPath, "\\")
	fName := splits[len(splits)-1]
	if len(splits) <= 1 {
		_, fName = path.Split(dicPath)
	}

	err = pdf.WritePdf(fName + ".pdf")
	if err != nil {
		panic(err)
	}
}

func PdfCmd() {
	var (
		w   float64
		h   float64
		pt  string
		dic string
	)

	flag.Float64Var(&h, "h", 0.0, "页面高度")
	flag.Float64Var(&w, "w", 0.0, "页面宽度")
	flag.StringVar(&pt, "p", "", "需要转换的文件夹集的目录，未设置时不生效")
	flag.StringVar(&dic, "d", "", "需要转换的单个文件夹，未设置时不生效")

	flag.Parse()

	if h == 0.0 || w == 0.0 {
		fmt.Println("未设置高度和宽度")
	}

	if dic != "" {
		fmt.Println("正在转换" + dic + "……")
		PdfImages(dic, Size{
			Weight: w,
			Height: h,
		})
	}

	if pt != "" {
		dir, err := ioutil.ReadDir(pt)
		if err != nil {
			panic(err)
		}
		for _, file := range dir {
			if !file.IsDir() {
				continue
			}

			fmt.Println("正在转换" + file.Name() + "……")
			PdfImages(pt+file.Name(), Size{
				Weight: w,
				Height: h,
			})
		}
	}

}
