package main

import (
	"fmt"
	"os"

	"github.com/binaryshogun/pdftoimage/pkg/image/writer"
	"github.com/binaryshogun/pdftoimage/pkg/pdf/converter"
	"github.com/binaryshogun/pdftoimage/pkg/pdf/scanner"

	"github.com/spf13/cobra"
)

var (
	pdfPath   string
	outputDir string
	format    string
)

var rootCmd = &cobra.Command{
	Use:   "pdftoimage",
	Short: "A tool to convert PDF file to images",
	Long:  `A tool to convert PDF file to images in various formats.`,
	Run: func(cmd *cobra.Command, args []string) {
		if pdfPath == "" {
			fmt.Println("PDF file path is required")
			_ = cmd.Usage()
			os.Exit(1)
		}

		var (
			s = scanner.NewScanner()
			w = writer.NewWriter()
			c = converter.NewConverter(s, w)
		)

		if err := c.Convert(pdfPath, outputDir, format); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error converting PDF to images: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("PDF conversion completed successfully!")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&pdfPath, "pdf", "p", "", "Path to the PDF file")
	rootCmd.PersistentFlags().StringVarP(&outputDir, "out", "o", "./images", "Output directory for images")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "jpeg", "Output image format (jpeg, png, gif)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
