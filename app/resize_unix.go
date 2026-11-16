// SPDX-License-Identifier: Unlicense OR MIT

//go:build (linux && !android) || freebsd || openbsd

package app

import (
	"gioui.org/io/pointer"
	"gioui.org/unit"
)

// resizeBorderDp is the client-side hit strip on undecorated windows.
// The window manager draws no frame, so the driver starts an interactive
// resize from this strip. The minimum size is unchanged; the compositor
// still clamps to the size hints.
const resizeBorderDp = unit.Dp(8)

// resizeDir is an edge or corner of an undecorated window.
type resizeDir int

const (
	resizeNone resizeDir = iota
	resizeTopLeft
	resizeTop
	resizeTopRight
	resizeRight
	resizeBottomRight
	resizeBottom
	resizeBottomLeft
	resizeLeft
)

// hitResize reports which edge contains the pointer. Coordinates and border
// are in the same pixel space as the window size.
func hitResize(x, y, width, height, border int) resizeDir {
	if border <= 0 || width <= 0 || height <= 0 {
		return resizeNone
	}
	north := y < border
	south := y >= height-border
	west := x < border
	east := x >= width-border
	switch {
	case north && west:
		return resizeTopLeft
	case north && east:
		return resizeTopRight
	case south && west:
		return resizeBottomLeft
	case south && east:
		return resizeBottomRight
	case north:
		return resizeTop
	case south:
		return resizeBottom
	case west:
		return resizeLeft
	case east:
		return resizeRight
	default:
		return resizeNone
	}
}

func (d resizeDir) cursor() pointer.Cursor {
	switch d {
	case resizeTopLeft:
		return pointer.CursorNorthWestResize
	case resizeTop:
		return pointer.CursorNorthResize
	case resizeTopRight:
		return pointer.CursorNorthEastResize
	case resizeRight:
		return pointer.CursorEastResize
	case resizeBottomRight:
		return pointer.CursorSouthEastResize
	case resizeBottom:
		return pointer.CursorSouthResize
	case resizeBottomLeft:
		return pointer.CursorSouthWestResize
	case resizeLeft:
		return pointer.CursorWestResize
	default:
		return pointer.CursorDefault
	}
}

// ewmhDir is the _NET_WM_MOVERESIZE direction for this edge.
func (d resizeDir) ewmhDir() int {
	switch d {
	case resizeTopLeft:
		return 0
	case resizeTop:
		return 1
	case resizeTopRight:
		return 2
	case resizeRight:
		return 3
	case resizeBottomRight:
		return 4
	case resizeBottom:
		return 5
	case resizeBottomLeft:
		return 6
	case resizeLeft:
		return 7
	default:
		return -1
	}
}
