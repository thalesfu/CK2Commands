package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡齐纳KatsinaBarony struct {
	feud.BaseBarony
}

var BKatsina卡齐纳 feud.Barony = &卡齐纳KatsinaBarony{}

func init() {
	f := BKatsina卡齐纳.(*卡齐纳KatsinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katsina",
		TitleName: "卡齐纳",
		TitleCode: "b_katsina",
	}
}
