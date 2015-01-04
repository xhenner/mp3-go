# id3

MP3 Indentification library for Go.

# Install

The platform ($GOROOT/bin) "go get" tool is the best method to install.

    go get github.com/xhenner/mp3-go

This downloads and installs the package into your $GOPATH. If you only want to
recompile, use "go install".

    go install github.com/xhenner/mp3-go

# Usage

An import allows access to the package.

    import (
        id3 "github.com/xhenner/mp3-go"
    )

# Quick Start

To access the information of a file, simply use 
    
    mp3File, err := mp3.Examine(path, slow)

if slow is true, the program will scan the whole file, if not, just the 500
first frames. The full scan is more precise (especially for VBR files) but
much slower
