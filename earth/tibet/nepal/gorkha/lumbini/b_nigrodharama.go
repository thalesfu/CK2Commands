package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼拘陀园NigrodharamaBarony struct {
	feud.BaseBarony
}

var BNigrodharama尼拘陀园 feud.Barony = &尼拘陀园NigrodharamaBarony{}

func init() {
    f := BNigrodharama尼拘陀园.(*尼拘陀园NigrodharamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nigrodharama",
		TitleName: "尼拘陀园",
		TitleCode: "b_nigrodharama",
	}
}
