package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢科姆利LukomlBarony struct {
	feud.BaseBarony
}

var BLukoml卢科姆利 feud.Barony = &卢科姆利LukomlBarony{}

func init() {
	f := BLukoml卢科姆利.(*卢科姆利LukomlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lukoml",
		TitleName: "卢科姆利",
		TitleCode: "b_lukoml",
	}
}
