package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨伊SayBarony struct {
	feud.BaseBarony
}

var BSay萨伊 feud.Barony = &萨伊SayBarony{}

func init() {
    f := BSay萨伊.(*萨伊SayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "say",
		TitleName: "萨伊",
		TitleCode: "b_say",
	}
}
