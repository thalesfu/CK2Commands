package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 龙城LuutBarony struct {
	feud.BaseBarony
}

var BLuut龙城 feud.Barony = &龙城LuutBarony{}

func init() {
    f := BLuut龙城.(*龙城LuutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luut",
		TitleName: "龙城",
		TitleCode: "b_luut",
	}
}
