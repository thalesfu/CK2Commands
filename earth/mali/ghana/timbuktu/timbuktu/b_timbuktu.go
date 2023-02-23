package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷巴克图TimbuktuBarony struct {
	feud.BaseBarony
}

var BTimbuktu廷巴克图 feud.Barony = &廷巴克图TimbuktuBarony{}

func init() {
	f := BTimbuktu廷巴克图.(*廷巴克图TimbuktuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timbuktu",
		TitleName: "廷巴克图",
		TitleCode: "b_timbuktu",
	}
}
