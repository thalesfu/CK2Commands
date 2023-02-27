package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨吉罗SatkhiraBarony struct {
	feud.BaseBarony
}

var BSatkhira萨吉罗 feud.Barony = &萨吉罗SatkhiraBarony{}

func init() {
    f := BSatkhira萨吉罗.(*萨吉罗SatkhiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satkhira",
		TitleName: "萨吉罗",
		TitleCode: "b_satkhira",
	}
}
