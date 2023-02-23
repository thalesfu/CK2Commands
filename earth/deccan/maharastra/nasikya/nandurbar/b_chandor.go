package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚多尔ChandorBarony struct {
	feud.BaseBarony
}

var BChandor尚多尔 feud.Barony = &尚多尔ChandorBarony{}

func init() {
	f := BChandor尚多尔.(*尚多尔ChandorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandor",
		TitleName: "尚多尔",
		TitleCode: "b_chandor",
	}
}
