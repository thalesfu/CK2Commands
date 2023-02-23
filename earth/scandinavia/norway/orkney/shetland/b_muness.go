package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆尼斯MunessBarony struct {
	feud.BaseBarony
}

var BMuness穆尼斯 feud.Barony = &穆尼斯MunessBarony{}

func init() {
	f := BMuness穆尼斯.(*穆尼斯MunessBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muness",
		TitleName: "穆尼斯",
		TitleCode: "b_muness",
	}
}
