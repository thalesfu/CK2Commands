package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都灵TorinoBarony struct {
	feud.BaseBarony
}

var BTorino都灵 feud.Barony = &都灵TorinoBarony{}

func init() {
	f := BTorino都灵.(*都灵TorinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torino",
		TitleName: "都灵",
		TitleCode: "b_torino",
	}
}
