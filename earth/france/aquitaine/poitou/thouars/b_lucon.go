package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕松LuconBarony struct {
	feud.BaseBarony
}

var BLucon吕松 feud.Barony = &吕松LuconBarony{}

func init() {
	f := BLucon吕松.(*吕松LuconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucon",
		TitleName: "吕松",
		TitleCode: "b_lucon",
	}
}
