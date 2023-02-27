package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涌波YurbaBarony struct {
	feud.BaseBarony
}

var BYurba涌波 feud.Barony = &涌波YurbaBarony{}

func init() {
    f := BYurba涌波.(*涌波YurbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yurba",
		TitleName: "涌波",
		TitleCode: "b_yurba",
	}
}
