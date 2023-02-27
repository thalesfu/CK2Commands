package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱索美奈KlazomeanaiBarony struct {
	feud.BaseBarony
}

var BKlazomeanai克莱索美奈 feud.Barony = &克莱索美奈KlazomeanaiBarony{}

func init() {
    f := BKlazomeanai克莱索美奈.(*克莱索美奈KlazomeanaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klazomeanai",
		TitleName: "克莱索美奈",
		TitleCode: "b_klazomeanai",
	}
}
