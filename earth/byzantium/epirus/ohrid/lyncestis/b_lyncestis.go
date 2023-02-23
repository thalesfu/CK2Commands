package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕恩盖斯提斯LyncestisBarony struct {
	feud.BaseBarony
}

var BLyncestis吕恩盖斯提斯 feud.Barony = &吕恩盖斯提斯LyncestisBarony{}

func init() {
	f := BLyncestis吕恩盖斯提斯.(*吕恩盖斯提斯LyncestisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyncestis",
		TitleName: "吕恩盖斯提斯",
		TitleCode: "b_lyncestis",
	}
}
