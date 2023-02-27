package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷奇科沃LychkovoBarony struct {
	feud.BaseBarony
}

var BLychkovo雷奇科沃 feud.Barony = &雷奇科沃LychkovoBarony{}

func init() {
    f := BLychkovo雷奇科沃.(*雷奇科沃LychkovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lychkovo",
		TitleName: "雷奇科沃",
		TitleCode: "b_lychkovo",
	}
}
