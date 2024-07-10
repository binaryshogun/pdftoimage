# PDF to Image Converter CLI Tool

A command-line interface (CLI) tool written in Go for converting PDF files to images of various formats.

[godoc.org/github.com/binaryshogun/pdftoimage](https://godoc.org/github.com/binaryshogun/pdftoimage) 
[goreportcard.com/report/github.com/github.com/binaryshogun/pdftoimage](https://goreportcard.com/report/github.com/binaryshogun/pdftoimage)

## Features

- Converts PDF files to images (PNG, JPEG)
- Easy to use with a simple command-line interface

## Installation

### Prerequisites

- Go programming language installed. [Download Go](https://golang.org/dl/).

### Installation

To install the CLI tool, run the following command:

```bash
go install github.com/binaryshogun/pdftoimage
```

## Usage

```bash
pdftoimage -p input.pdf --format png -o images/
```

### Options

- `--pdf` or `-p`: (required) Path to the PDF file.
- `input.pdf`: Path to the input PDF file.
- `--format` or `-f`:  Output image format (`png`, `jpeg`, `jpg`). Default to `jpeg`.
- `--out` or `-o`: Output directory for images. Default to local `images` folder.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/binaryshogun/pdftoimage/LICENCE.md) file for details.
