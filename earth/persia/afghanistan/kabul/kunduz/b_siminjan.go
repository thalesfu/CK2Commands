package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悉泯健SiminjanBarony struct {
	feud.BaseBarony
}

var BSiminjan悉泯健 feud.Barony = &悉泯健SiminjanBarony{}

func init() {
	f := BSiminjan悉泯健.(*悉泯健SiminjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siminjan",
		TitleName: "悉泯健",
		TitleCode: "b_siminjan",
	}
}
