package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧卡兹Al_ukazBarony struct {
	feud.BaseBarony
}

var BAl_ukaz欧卡兹 feud.Barony = &欧卡兹Al_ukazBarony{}

func init() {
    f := BAl_ukaz欧卡兹.(*欧卡兹Al_ukazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_ukaz",
		TitleName: "欧卡兹",
		TitleCode: "b_al_ukaz",
	}
}
