package cmd

import (
	"fmt"
	"os"

	"github.com/binaryshogun/pdftoimage/pkg/converter"
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

		c := converter.NewConverter()

		if err := c.Convert(pdfPath, outputDir, format); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error converting PDF to images: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("PDF conversion completed successfully!")
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&pdfPath, "pdf", "p", "", "Path to the PDF file")
	rootCmd.PersistentFlags().StringVarP(&outputDir, "out", "o", "./images", "Output directory for images")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "jpeg", "Output image format (jpeg, png, gif)")
}
