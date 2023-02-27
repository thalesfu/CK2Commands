package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫罗姆格拉HromglaBarony struct {
	feud.BaseBarony
}

var BHromgla赫罗姆格拉 feud.Barony = &赫罗姆格拉HromglaBarony{}

func init() {
    f := BHromgla赫罗姆格拉.(*赫罗姆格拉HromglaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hromgla",
		TitleName: "赫罗姆格拉",
		TitleCode: "b_hromgla",
	}
}
