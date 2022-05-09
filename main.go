package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/jung-kurt/gofpdf"
	"github.com/nfnt/resize"
)

var (
	ratio = flag.Float64("ratio", 0.7, "set ratio")
)

func main() {
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		name := flag.Arg(i)
		err := extractPdf(name)
		if err != nil {
			log.Fatalf("shrink %q failed: %v", name, err)
		}
	}
}

func extractPdf(name string) error {
	doc, err := fitz.New(name)
	if err != nil {
		return err
	}
	defer doc.Close()

	name = filepath.Base(name)
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)

	// Ensure pdf file has no text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			return err
		}

		text = strings.TrimSpace(text)
		if text != "" {
			return fmt.Errorf("page %d has text", n+1)
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			return err
		}

		rc := img.Bounds()
		w := rc.Dx()
		h := rc.Dy()

		fmt.Printf("page %d: image: %dx%d\n", n+1, w, h)

		newImage := resize.Resize(uint(float64(w)*float64(*ratio)), 0, img, resize.Lanczos3)
		imgName := fmt.Sprintf("%s%03d.jpg", base, n)

		f, err := os.Create(imgName)
		if err != nil {
			return err
		}

		err = jpeg.Encode(f, newImage, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			f.Close()
			return err
		}
		f.Close()

		pageW, pageH := pdf.GetPageSize()

		pdf.Image(imgName, 0, 0, pageW, pageH, false, "", 0, "")
	}

	pdfName := fmt.Sprintf("%s.pdf", base)
	return pdf.OutputFileAndClose(pdfName)
}
