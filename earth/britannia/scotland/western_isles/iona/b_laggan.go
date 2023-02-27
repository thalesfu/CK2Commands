package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉根LagganBarony struct {
	feud.BaseBarony
}

var BLaggan拉根 feud.Barony = &拉根LagganBarony{}

func init() {
    f := BLaggan拉根.(*拉根LagganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laggan",
		TitleName: "拉根",
		TitleCode: "b_laggan",
	}
}
