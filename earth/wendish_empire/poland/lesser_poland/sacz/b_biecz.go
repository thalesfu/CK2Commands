package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别奇BieczBarony struct {
	feud.BaseBarony
}

var BBiecz别奇 feud.Barony = &别奇BieczBarony{}

func init() {
    f := BBiecz别奇.(*别奇BieczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biecz",
		TitleName: "别奇",
		TitleCode: "b_biecz",
	}
}
