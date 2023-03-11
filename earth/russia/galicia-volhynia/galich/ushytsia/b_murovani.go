package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆罗瓦尼MurovaniBarony struct {
	feud.BaseBarony
}

var BMurovani穆罗瓦尼 feud.Barony = &穆罗瓦尼MurovaniBarony{}

func init() {
    f := BMurovani穆罗瓦尼.(*穆罗瓦尼MurovaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murovani",
		TitleName: "穆罗瓦尼",
		TitleCode: "b_murovani",
	}
}
