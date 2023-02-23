package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗申RojhanBarony struct {
	feud.BaseBarony
}

var BRojhan罗申 feud.Barony = &罗申RojhanBarony{}

func init() {
	f := BRojhan罗申.(*罗申RojhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rojhan",
		TitleName: "罗申",
		TitleCode: "b_rojhan",
	}
}
