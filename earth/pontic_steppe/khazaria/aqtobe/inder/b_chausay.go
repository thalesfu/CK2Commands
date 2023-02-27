package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔_赛ChausayBarony struct {
	feud.BaseBarony
}

var BChausay乔_赛 feud.Barony = &乔_赛ChausayBarony{}

func init() {
    f := BChausay乔_赛.(*乔_赛ChausayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chausay",
		TitleName: "乔_赛",
		TitleCode: "b_chausay",
	}
}
