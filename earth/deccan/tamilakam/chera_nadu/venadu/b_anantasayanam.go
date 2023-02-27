package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿难多娑耶南摩AnantasayanamBarony struct {
	feud.BaseBarony
}

var BAnantasayanam阿难多娑耶南摩 feud.Barony = &阿难多娑耶南摩AnantasayanamBarony{}

func init() {
    f := BAnantasayanam阿难多娑耶南摩.(*阿难多娑耶南摩AnantasayanamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anantasayanam",
		TitleName: "阿难多娑耶南摩",
		TitleCode: "b_anantasayanam",
	}
}
