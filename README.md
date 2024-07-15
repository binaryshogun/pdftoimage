# PDF to Image Converter CLI Tool

A command-line interface (CLI) tool written in Go for converting PDF files to images of various formats.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/binaryshogun/pdftoimage)
![example workflow](https://github.com/github/docs/actions/workflows/go.yml/badge.svg)

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
