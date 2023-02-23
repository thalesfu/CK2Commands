package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕迪尔马TendirmaBarony struct {
	feud.BaseBarony
}

var BTendirma滕迪尔马 feud.Barony = &滕迪尔马TendirmaBarony{}

func init() {
	f := BTendirma滕迪尔马.(*滕迪尔马TendirmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tendirma",
		TitleName: "滕迪尔马",
		TitleCode: "b_tendirma",
	}
}
