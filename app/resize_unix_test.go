// SPDX-License-Identifier: Unlicense OR MIT

//go:build (linux && !android) || freebsd || openbsd

package app

import "testing"

func TestHitResize(t *testing.T) {
	const (
		w, h = 800, 544
		b    = 8
	)
	cases := []struct {
		x, y int
		want resizeDir
	}{
		{0, 0, resizeTopLeft},
		{7, 3, resizeTopLeft},
		{400, 0, resizeTop},
		{799, 0, resizeTopRight},
		{799, 200, resizeRight},
		{799, 543, resizeBottomRight},
		{400, 540, resizeBottom},
		{0, 543, resizeBottomLeft},
		{0, 200, resizeLeft},
		{8, 8, resizeNone},
		{400, 272, resizeNone},
		{792, 8, resizeRight},
	}
	for _, c := range cases {
		if got := hitResize(c.x, c.y, w, h, b); got != c.want {
			t.Errorf("hitResize(%d,%d) = %v, want %v", c.x, c.y, got, c.want)
		}
	}
	if hitResize(0, 0, w, h, 0) != resizeNone {
		t.Fatal("zero border should not resize")
	}
	if hitResize(0, 0, 0, h, b) != resizeNone {
		t.Fatal("empty window should not resize")
	}
}

func TestResizeEWMHDirs(t *testing.T) {
	got := []int{
		resizeTopLeft.ewmhDir(),
		resizeTop.ewmhDir(),
		resizeTopRight.ewmhDir(),
		resizeRight.ewmhDir(),
		resizeBottomRight.ewmhDir(),
		resizeBottom.ewmhDir(),
		resizeBottomLeft.ewmhDir(),
		resizeLeft.ewmhDir(),
	}
	for i, dir := range got {
		if dir != i {
			t.Fatalf("dir %d = %d", i, dir)
		}
	}
	if resizeNone.ewmhDir() != -1 {
		t.Fatal("none has no EWMH direction")
	}
}
