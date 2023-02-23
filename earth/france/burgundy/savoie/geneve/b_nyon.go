package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼翁NyonBarony struct {
	feud.BaseBarony
}

var BNyon尼翁 feud.Barony = &尼翁NyonBarony{}

func init() {
	f := BNyon尼翁.(*尼翁NyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyon",
		TitleName: "尼翁",
		TitleCode: "b_nyon",
	}
}
