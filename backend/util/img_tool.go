package util

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func TurnImgToTransparent(path string) []byte {
	cwd, _ := os.Getwd()
	log.Printf("当前工作目录：%s", cwd)
	maskFile, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("读取蒙版失败：%v", err))
	}
	defer maskFile.Close()

	// 解码图片
	maskImg, _, err := image.Decode(maskFile)
	if err != nil {
		panic(fmt.Sprintf("解码蒙版失败：%v", err))
	}

	//  核心操作：只保留 Alpha 通道，RGB 设为透明
	rgba := image.NewRGBA(maskImg.Bounds())
	draw.Draw(rgba, rgba.Bounds(), maskImg, image.Point{}, draw.Src)

	// 遍历所有像素，清空 RGB，保留 Alpha
	bounds := rgba.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 获取像素的 RGBA 值（a 范围是 0-65535）
			_, _, _, a := rgba.At(x, y).RGBA()
			if a == 0 {
				rgba.SetRGBA(x, y, color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 0, // 仅保留 Alpha 通道
				})
				continue
			}
			rgba.SetRGBA(x, y, color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 1, // 人眼不可见的伪透明色。
			})
		}
	}

	//  重新编码为 PNG 二进制
	var maskBytes bytes.Buffer
	if err := png.Encode(&maskBytes, rgba); err != nil {
		panic(fmt.Sprintf("重编码蒙版失败：%v", err))
	}
	maskImageBytes := maskBytes.Bytes()
	return maskImageBytes
}
