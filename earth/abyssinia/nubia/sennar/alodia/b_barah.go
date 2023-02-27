package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉BarahBarony struct {
	feud.BaseBarony
}

var BBarah巴拉 feud.Barony = &巴拉BarahBarony{}

func init() {
    f := BBarah巴拉.(*巴拉BarahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barah",
		TitleName: "巴拉",
		TitleCode: "b_barah",
	}
}
