package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍尔甘SergenBarony struct {
	feud.BaseBarony
}

var BSergen舍尔甘 feud.Barony = &舍尔甘SergenBarony{}

func init() {
	f := BSergen舍尔甘.(*舍尔甘SergenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sergen",
		TitleName: "舍尔甘",
		TitleCode: "b_sergen",
	}
}
