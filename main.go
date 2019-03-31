package main

import (
	"flag"
	"fmt"
	"github.com/rakyll/command"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

type GifCommand struct {
	qualityFlag *int
	delayFlag   *int
}

type CompressCommand struct {
	qualityFlag *int
}

type DecompressCommand struct {
}

func (cmd *GifCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	cmd.qualityFlag = fs.Int("quality", 10, "quadtree depth (more is better quality)")
	cmd.delayFlag = fs.Int("delayMS", 500, "frame delay in ms")
	return fs
}

func (cmd *CompressCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	cmd.qualityFlag = fs.Int("quality", -1, "quadtree depth (more is better quality, -1 is unbounded/max)")
	return fs
}

func (cmd *DecompressCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	return fs
}

func (cmd *GifCommand) Run(args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <inpath> <outpath>\n", os.Args[0])
		os.Exit(0)
	}

	inPath := args[0]
	outPath := args[1]

	inImg, err := LoadImage(inPath)
	if err != nil {
		panic(err)
	}

	qt, err := BuildQuadTree(inImg)
	if err != nil {
		panic(err)
	}

	outGif := &gif.GIF{}
	for i := 1; i < *cmd.qualityFlag; i++ {
		compressedQ, _ := qt.Compress(i)
		img := compressedQ.ToImage()

		pImg := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.Draw(pImg, pImg.Rect, img, img.Bounds().Min, draw.Over)

		outGif.Image = append(outGif.Image, pImg)
		outGif.Delay = append(outGif.Delay, *cmd.delayFlag/10)
	}

	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gif.EncodeAll(f, outGif)
}

func (cmd *CompressCommand) Run(args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <inpath> <outpath>\n", os.Args[0])
		os.Exit(0)
	}

	inPath := args[0]
	outPath := args[1]

	inImg, err := LoadImage(inPath)
	if err != nil {
		panic(err)
	}

	qt, err := BuildQuadTree(inImg)
	if err != nil {
		panic(err)
	}

	qt, err = qt.Compress(*cmd.qualityFlag)
	if err != nil {
		panic(err)
	}

	err = qt.SerializeToFile(outPath)
	if err != nil {
		panic(err)
	}
}

func (cmd *DecompressCommand) Run(args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <inpath> <outpath>\n", os.Args[0])
		os.Exit(0)
	}

	inPath := args[0]
	outPath := args[1]

	inQt, err := LoadQuadTreeFromFile(inPath)
	if err != nil {
		panic(err)
	}

	img := inQt.ToImage()

	err = WriteImage(img, outPath)
	if err != nil {
		panic(err)
	}
}

func main() {
	command.On("gif", "generate a gif", &GifCommand{}, []string{})
	command.On("compress", "compress to .quadtree file", &CompressCommand{}, []string{})
	command.On("decompress", "decompress .quadtree file", &DecompressCommand{}, []string{})

	command.Parse()
	command.Run()
}
