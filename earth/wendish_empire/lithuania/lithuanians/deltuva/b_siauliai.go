package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希奥利艾SiauliaiBarony struct {
	feud.BaseBarony
}

var BSiauliai希奥利艾 feud.Barony = &希奥利艾SiauliaiBarony{}

func init() {
    f := BSiauliai希奥利艾.(*希奥利艾SiauliaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siauliai",
		TitleName: "希奥利艾",
		TitleCode: "b_siauliai",
	}
}
