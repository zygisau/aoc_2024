package day9

import (
	"strconv"
	"testing"
)

func TestDecode(t *testing.T) {
	a := "2333133121414131402"
	want := "00...111...2...333.44.5555.6666.777.888899"

	out, _, err := DecodeLine(a)
	if err != nil {
		t.Error(err)
	}

	outStringified := ""
	for _, val := range out {
		if val == -1 {
			outStringified += "."
		} else {
			outStringified += strconv.Itoa(val)
		}
	}

	if want != outStringified {
		t.Fatalf("TestDecode()\nwant: %s\n got: %v", want, outStringified)
	}
}

func TestReformat(t *testing.T) {
	a := "2333133121414131402"
	want := "0099811188827773336446555566.............."
	disk, freeSpaceIdxs, err := DecodeLine(a)
	if err != nil {
		t.Error(err)
	}
	out := ReformatDisk(disk, freeSpaceIdxs)

	outStringified := ""
	for _, val := range out {
		if val == -1 {
			outStringified += "."
		} else {
			outStringified += strconv.Itoa(val)
		}
	}

	if want != outStringified {
		t.Fatalf("TestReformat()\nwant: %s\n got: %v", want, outStringified)
	}
}

func TestCalcChecksum(t *testing.T) {
	a := "2333133121414131402"
	want := uint64(1928)
	disk, freeSpaceIdxs, err := DecodeLine(a)
	if err != nil {
		t.Error(err)
	}
	reformatted := ReformatDisk(disk, freeSpaceIdxs)
	out := CalcChecksum(reformatted)

	if want != out {
		t.Fatalf("TestCalcChecksum()\nwant: %d\n got: %d", want, out)
	}
}

func TestReadAndCalcChecksum(t *testing.T) {
	filename := "input.txt"
	a, err := ReadLine(filename)
	if err != nil {
		t.Error(err)
	}

	want := uint64(6346871685398)
	disk, freeSpaceIdxs, err := DecodeLine(a)
	if err != nil {
		t.Error(err)
	}

	reformatted := ReformatDisk(disk, freeSpaceIdxs)
	out := CalcChecksum(reformatted)

	if want != out {
		t.Fatalf("TestReadAndCalcChecksum()\nwant: %d\n got: %d", want, out)
	}
}
