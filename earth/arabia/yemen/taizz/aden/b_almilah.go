package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米拉AlmilahBarony struct {
	feud.BaseBarony
}

var BAlmilah米拉 feud.Barony = &米拉AlmilahBarony{}

func init() {
	f := BAlmilah米拉.(*米拉AlmilahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almilah",
		TitleName: "米拉",
		TitleCode: "b_almilah",
	}
}
