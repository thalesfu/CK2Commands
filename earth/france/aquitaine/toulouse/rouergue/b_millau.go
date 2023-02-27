package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米约MillauBarony struct {
	feud.BaseBarony
}

var BMillau米约 feud.Barony = &米约MillauBarony{}

func init() {
    f := BMillau米约.(*米约MillauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "millau",
		TitleName: "米约",
		TitleCode: "b_millau",
	}
}
