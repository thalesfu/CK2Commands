package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢斯LuusBarony struct {
	feud.BaseBarony
}

var BLuus卢斯 feud.Barony = &卢斯LuusBarony{}

func init() {
    f := BLuus卢斯.(*卢斯LuusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luus",
		TitleName: "卢斯",
		TitleCode: "b_luus",
	}
}
