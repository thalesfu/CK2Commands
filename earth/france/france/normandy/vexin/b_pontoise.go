package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬图瓦兹PontoiseBarony struct {
	feud.BaseBarony
}

var BPontoise蓬图瓦兹 feud.Barony = &蓬图瓦兹PontoiseBarony{}

func init() {
    f := BPontoise蓬图瓦兹.(*蓬图瓦兹PontoiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontoise",
		TitleName: "蓬图瓦兹",
		TitleCode: "b_pontoise",
	}
}
