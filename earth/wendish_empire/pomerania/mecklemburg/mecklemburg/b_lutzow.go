package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕措LutzowBarony struct {
	feud.BaseBarony
}

var BLutzow吕措 feud.Barony = &吕措LutzowBarony{}

func init() {
    f := BLutzow吕措.(*吕措LutzowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lutzow",
		TitleName: "吕措",
		TitleCode: "b_lutzow",
	}
}
