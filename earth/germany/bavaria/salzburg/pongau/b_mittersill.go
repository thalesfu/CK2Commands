package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米特西尔MittersillBarony struct {
	feud.BaseBarony
}

var BMittersill米特西尔 feud.Barony = &米特西尔MittersillBarony{}

func init() {
	f := BMittersill米特西尔.(*米特西尔MittersillBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mittersill",
		TitleName: "米特西尔",
		TitleCode: "b_mittersill",
	}
}
