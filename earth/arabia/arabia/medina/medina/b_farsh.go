package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉斯FarshBarony struct {
	feud.BaseBarony
}

var BFarsh法拉斯 feud.Barony = &法拉斯FarshBarony{}

func init() {
	f := BFarsh法拉斯.(*法拉斯FarshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farsh",
		TitleName: "法拉斯",
		TitleCode: "b_farsh",
	}
}
