package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥登佩赫OdenpahBarony struct {
	feud.BaseBarony
}

var BOdenpah奥登佩赫 feud.Barony = &奥登佩赫OdenpahBarony{}

func init() {
	f := BOdenpah奥登佩赫.(*奥登佩赫OdenpahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odenpah",
		TitleName: "奥登佩赫",
		TitleCode: "b_odenpah",
	}
}
