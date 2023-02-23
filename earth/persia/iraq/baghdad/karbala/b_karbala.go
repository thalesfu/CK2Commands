package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔巴拉KarbalaBarony struct {
	feud.BaseBarony
}

var BKarbala卡尔巴拉 feud.Barony = &卡尔巴拉KarbalaBarony{}

func init() {
	f := BKarbala卡尔巴拉.(*卡尔巴拉KarbalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karbala",
		TitleName: "卡尔巴拉",
		TitleCode: "b_karbala",
	}
}
