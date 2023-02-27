package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢隆LoulonBarony struct {
	feud.BaseBarony
}

var BLoulon卢隆 feud.Barony = &卢隆LoulonBarony{}

func init() {
    f := BLoulon卢隆.(*卢隆LoulonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loulon",
		TitleName: "卢隆",
		TitleCode: "b_loulon",
	}
}
