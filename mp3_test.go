// Copyright 2014 Xavier Henner. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package mp3

import (
	"testing"
)

const (
	testFileC = "testcbr.mp3"
	testFileV = "testvbr.mp3"
)

func TestCBR(t *testing.T) {
	mp3File, err := Examine(testFileC, false)
	if err != nil {
		t.Errorf("Can't read file")
	}

	if mp3File.Version != "1" {
		t.Errorf("Can't detect Version")
	}

	if mp3File.Layer != "III" {
		t.Errorf("Can't detect Layer")
	}

	if mp3File.Mode != "Join Stereo" {
		t.Errorf("Can't detect Mode")
	}

	if mp3File.Bitrate != 128 {
		t.Errorf("Can't detect Bitrate")
	}

	if mp3File.Type != "CBR" {
		t.Errorf("Can't detect CBR")
	}

	if mp3File.Sampling != 48000 {
		t.Errorf("Can't detect Sampling")
	}

	if mp3File.Size != 1126528 {
		t.Errorf("Invalid size")
	}

	if int(mp3File.Length) != 70 {
		t.Errorf("Can't detect Length")
	}
}

func TestVBR(t *testing.T) {

	mp3File, err := Examine(testFileV, true)
	if err != nil {
		t.Errorf("Can't read file")
	}

	if mp3File.Version != "1" {
		t.Errorf("Can't detect Version")
	}

	if mp3File.Layer != "III" {
		t.Errorf("Can't detect Layer")
	}

	if mp3File.Mode != "Stereo" {
		t.Errorf("Can't detect Mode")
	}

	if mp3File.Bitrate != 192 {
		t.Errorf("Can't detect Bitrate")
	}

	if mp3File.Type != "VBR" {
		t.Errorf("Can't detect VBR")
	}

	if mp3File.Sampling != 48000 {
		t.Errorf("Can't detect Sampling")
	}

	if mp3File.Size != 1630336 {
		t.Errorf("Invalid size")

	}

	if int(mp3File.Length) != 70 {
		t.Errorf("Can't detect Length")
	}
}

func TestVBRfast(t *testing.T) {
	mp3File, err := Examine(testFileV, false)

	if err != nil {
		t.Errorf("Can't read file")
	}

	if mp3File.Length < 67 || mp3File.Length > 73 {
		t.Errorf("Can't detect Length")
	}

	if mp3File.Bitrate != 192 {
		t.Errorf("Can't detect Bitrate")
	}

	if mp3File.Type != "VBR" {
		t.Errorf("Can't detect VBR")
	}
}
