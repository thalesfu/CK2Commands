package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢顿LutonBarony struct {
	feud.BaseBarony
}

var BLuton卢顿 feud.Barony = &卢顿LutonBarony{}

func init() {
	f := BLuton卢顿.(*卢顿LutonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luton",
		TitleName: "卢顿",
		TitleCode: "b_luton",
	}
}
