package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳博姆利LubomlBarony struct {
	feud.BaseBarony
}

var BLuboml柳博姆利 feud.Barony = &柳博姆利LubomlBarony{}

func init() {
    f := BLuboml柳博姆利.(*柳博姆利LubomlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luboml",
		TitleName: "柳博姆利",
		TitleCode: "b_luboml",
	}
}
