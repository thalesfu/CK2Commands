package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢伊罗LuiroBarony struct {
	feud.BaseBarony
}

var BLuiro卢伊罗 feud.Barony = &卢伊罗LuiroBarony{}

func init() {
	f := BLuiro卢伊罗.(*卢伊罗LuiroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luiro",
		TitleName: "卢伊罗",
		TitleCode: "b_luiro",
	}
}
