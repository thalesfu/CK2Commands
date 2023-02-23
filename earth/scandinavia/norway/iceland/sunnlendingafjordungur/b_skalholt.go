package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯考尔霍特SkalholtBarony struct {
	feud.BaseBarony
}

var BSkalholt斯考尔霍特 feud.Barony = &斯考尔霍特SkalholtBarony{}

func init() {
	f := BSkalholt斯考尔霍特.(*斯考尔霍特SkalholtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skalholt",
		TitleName: "斯考尔霍特",
		TitleCode: "b_skalholt",
	}
}
