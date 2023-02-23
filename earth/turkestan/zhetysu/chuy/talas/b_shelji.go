package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希勒吉SheljiBarony struct {
	feud.BaseBarony
}

var BShelji希勒吉 feud.Barony = &希勒吉SheljiBarony{}

func init() {
	f := BShelji希勒吉.(*希勒吉SheljiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shelji",
		TitleName: "希勒吉",
		TitleCode: "b_shelji",
	}
}
