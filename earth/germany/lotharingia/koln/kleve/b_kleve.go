package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱沃KleveBarony struct {
	feud.BaseBarony
}

var BKleve克莱沃 feud.Barony = &克莱沃KleveBarony{}

func init() {
	f := BKleve克莱沃.(*克莱沃KleveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kleve",
		TitleName: "克莱沃",
		TitleCode: "b_kleve",
	}
}
