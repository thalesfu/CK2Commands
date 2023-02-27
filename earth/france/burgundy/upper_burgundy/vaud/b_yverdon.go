package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊韦尔东YverdonBarony struct {
	feud.BaseBarony
}

var BYverdon伊韦尔东 feud.Barony = &伊韦尔东YverdonBarony{}

func init() {
    f := BYverdon伊韦尔东.(*伊韦尔东YverdonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yverdon",
		TitleName: "伊韦尔东",
		TitleCode: "b_yverdon",
	}
}
