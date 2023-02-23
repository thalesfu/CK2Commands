package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕德斯堡LudersburgBarony struct {
	feud.BaseBarony
}

var BLudersburg吕德斯堡 feud.Barony = &吕德斯堡LudersburgBarony{}

func init() {
	f := BLudersburg吕德斯堡.(*吕德斯堡LudersburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ludersburg",
		TitleName: "吕德斯堡",
		TitleCode: "b_ludersburg",
	}
}
