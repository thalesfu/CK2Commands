package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊博拉IboraBarony struct {
	feud.BaseBarony
}

var BIbora伊博拉 feud.Barony = &伊博拉IboraBarony{}

func init() {
	f := BIbora伊博拉.(*伊博拉IboraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ibora",
		TitleName: "伊博拉",
		TitleCode: "b_ibora",
	}
}
