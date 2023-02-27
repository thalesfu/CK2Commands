package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕达霍TendahoBarony struct {
	feud.BaseBarony
}

var BTendaho滕达霍 feud.Barony = &滕达霍TendahoBarony{}

func init() {
    f := BTendaho滕达霍.(*滕达霍TendahoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tendaho",
		TitleName: "滕达霍",
		TitleCode: "b_tendaho",
	}
}
