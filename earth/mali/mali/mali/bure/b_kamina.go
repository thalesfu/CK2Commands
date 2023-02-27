package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡米纳KaminaBarony struct {
	feud.BaseBarony
}

var BKamina卡米纳 feud.Barony = &卡米纳KaminaBarony{}

func init() {
    f := BKamina卡米纳.(*卡米纳KaminaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamina",
		TitleName: "卡米纳",
		TitleCode: "b_kamina",
	}
}
