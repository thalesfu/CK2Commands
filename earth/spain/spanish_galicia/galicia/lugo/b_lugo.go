package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢戈LugoBarony struct {
	feud.BaseBarony
}

var BLugo卢戈 feud.Barony = &卢戈LugoBarony{}

func init() {
    f := BLugo卢戈.(*卢戈LugoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lugo",
		TitleName: "卢戈",
		TitleCode: "b_lugo",
	}
}
