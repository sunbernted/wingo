package main

import (
	"github.com/BurntSushi/wingo/focus"
)

func ignoreRootFocus(modeByte, detailByte byte) bool {
	mode, detail := focus.Modes[modeByte], focus.Details[detailByte]

	if mode == "NotifyGrab" || mode == "NotifyUngrab" {
		return true
	}
	if detail == "NotifyAncestor" ||
		detail == "NotifyInferior" ||
		detail == "NotifyVirtual" ||
		detail == "NotifyNonlinear" ||
		detail == "NotifyNonlinearVirtual" ||
		detail == "NotifyPointer" {

		return true
	}
	// Only accept modes: NotifyNormal and NotifyWhileGrabbed
	// Only accept details: NotifyPointerRoot, NotifyNone
	return false
}
