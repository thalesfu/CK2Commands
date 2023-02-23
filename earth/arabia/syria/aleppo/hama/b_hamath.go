package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马斯HamathBarony struct {
	feud.BaseBarony
}

var BHamath哈马斯 feud.Barony = &哈马斯HamathBarony{}

func init() {
	f := BHamath哈马斯.(*哈马斯HamathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamath",
		TitleName: "哈马斯",
		TitleCode: "b_hamath",
	}
}
