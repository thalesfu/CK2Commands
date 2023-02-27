package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔戈波尔KargopolBarony struct {
	feud.BaseBarony
}

var BKargopol卡尔戈波尔 feud.Barony = &卡尔戈波尔KargopolBarony{}

func init() {
    f := BKargopol卡尔戈波尔.(*卡尔戈波尔KargopolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kargopol",
		TitleName: "卡尔戈波尔",
		TitleCode: "b_kargopol",
	}
}
