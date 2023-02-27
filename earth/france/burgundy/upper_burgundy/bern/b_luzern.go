package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢塞恩LuzernBarony struct {
	feud.BaseBarony
}

var BLuzern卢塞恩 feud.Barony = &卢塞恩LuzernBarony{}

func init() {
    f := BLuzern卢塞恩.(*卢塞恩LuzernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luzern",
		TitleName: "卢塞恩",
		TitleCode: "b_luzern",
	}
}
