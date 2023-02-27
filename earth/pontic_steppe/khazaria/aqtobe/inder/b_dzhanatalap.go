package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾纳塔拉普DzhanatalapBarony struct {
	feud.BaseBarony
}

var BDzhanatalap贾纳塔拉普 feud.Barony = &贾纳塔拉普DzhanatalapBarony{}

func init() {
    f := BDzhanatalap贾纳塔拉普.(*贾纳塔拉普DzhanatalapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhanatalap",
		TitleName: "贾纳塔拉普",
		TitleCode: "b_dzhanatalap",
	}
}
