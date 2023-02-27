package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅那摩提MainamatiBarony struct {
	feud.BaseBarony
}

var BMainamati梅那摩提 feud.Barony = &梅那摩提MainamatiBarony{}

func init() {
    f := BMainamati梅那摩提.(*梅那摩提MainamatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mainamati",
		TitleName: "梅那摩提",
		TitleCode: "b_mainamati",
	}
}
