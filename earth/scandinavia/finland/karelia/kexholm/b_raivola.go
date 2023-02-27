package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖沃拉RaivolaBarony struct {
	feud.BaseBarony
}

var BRaivola赖沃拉 feud.Barony = &赖沃拉RaivolaBarony{}

func init() {
    f := BRaivola赖沃拉.(*赖沃拉RaivolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raivola",
		TitleName: "赖沃拉",
		TitleCode: "b_raivola",
	}
}
