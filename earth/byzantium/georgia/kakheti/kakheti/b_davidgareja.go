package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大卫加列贾DavidgarejaBarony struct {
	feud.BaseBarony
}

var BDavidgareja大卫加列贾 feud.Barony = &大卫加列贾DavidgarejaBarony{}

func init() {
    f := BDavidgareja大卫加列贾.(*大卫加列贾DavidgarejaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "davidgareja",
		TitleName: "大卫加列贾",
		TitleCode: "b_davidgareja",
	}
}
