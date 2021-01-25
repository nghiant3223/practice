package main

import (
	"math"

	"github.com/fogleman/gg"
)

const (
	outPath string = "out.png"

	lineWidth   float64 = 1
	trunkHeight float64 = 512
	width       float64 = 1920
	height      float64 = 1080

	degenerationAngle   float64 = math.Pi / 6
	degenerationLength  float64 = 0.5
	minimumBranchLength float64 = 5

	colorRed   = 0
	colorBlue  = 0
	colorGreen = 0
)

func main() {
	ctx := prepare()
	drawTree(ctx)
	saveTree(ctx)
}

func saveTree(ctx *gg.Context) {
	ctx.SavePNG(outPath)
}

func drawTree(ctx *gg.Context) {
	drawMiddleTrunk(ctx)
	drawBranch(ctx, width/2, height-trunkHeight, trunkHeight)
}

func prepare() *gg.Context {
	intWidth, intHeight := int(width), int(height)
	ctx := gg.NewContext(intWidth, intHeight)
	ctx.SetRGB(colorRed, colorBlue, colorGreen)
	ctx.SetLineWidth(lineWidth)
	return ctx
}

func drawLine(ctx *gg.Context, x1, y1, x2, y2 float64) {
	ctx.Push()
	ctx.DrawLine(x1, y1, x2, y2)
	ctx.Stroke()
}

func drawMiddleTrunk(ctx *gg.Context) {
	x := width / 2
	y1 := height
	y2 := height - trunkHeight
	drawLine(ctx, x, y1, x, y2)
}

func drawBranch(ctx *gg.Context, px, py float64, pLen float64) {
	branchLength := pLen * degenerationLength
	if branchLength <= minimumBranchLength {
		return
	}

	cos := math.Cos(degenerationAngle)
	sin := math.Sin(degenerationAngle)
	minusCos := math.Cos(-degenerationAngle)
	minusSin := math.Sin(-degenerationAngle)

	leftBranchY := py - branchLength
	leftBranchX := px
	leftBranchX = cos*leftBranchX - sin*leftBranchY
	leftBranchY = sin*leftBranchX + cos*leftBranchY

	rightBranchY := py - branchLength
	rightBranchX := px
	rightBranchX = minusCos*rightBranchX - minusSin*rightBranchY
	rightBranchY = minusSin*rightBranchX + minusCos*rightBranchY

	drawLine(ctx, px, py, leftBranchX, leftBranchY)
	drawLine(ctx, px, py, rightBranchX, rightBranchY)

	drawBranch(ctx, leftBranchX, leftBranchY, branchLength)
	drawBranch(ctx, rightBranchX, rightBranchY, branchLength)
}
