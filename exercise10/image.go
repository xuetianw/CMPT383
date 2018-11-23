package exercise10

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func DrawCircle(outerRadius, innerRadius int, outputFile string) {
	bounds := image.Rectangle{
		Min:image.Point{0, 0}, Max:image.Point{200,200},
	}

	backGround := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}  //  R, G, B, Alpha
	black := color.RGBA{0, 0, 0, 255}  //  R, G, B, Alpha

	draw.Draw(backGround, backGround.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	new_png_file := outputFile

	x, y := 0, 0
	for y = 100-outerRadius; y <= 100+outerRadius; y++{
		for x = 100-outerRadius; x<= 100+outerRadius; x++{
			if ((x - 100) * (x - 100) + (y - 100) *  (y - 100)) < (outerRadius * outerRadius) {
				backGround.Set(x, y, black)
			}
		}
	}

	for y = 100-innerRadius; y <= 100+innerRadius; y++{
		for x = 100-innerRadius; x<= 100+innerRadius; x++{
			if ((x - 100) * (x - 100) + (y - 100) *  (y - 100)) < (innerRadius * innerRadius) {
				backGround.Set(x, y, white)
			}
		}
	}

	myfile, _ := os.Create(new_png_file)
	png.Encode(myfile, backGround)
}

//func main ()  {
//	DrawCircle(40, 20, "out.png")
//}
