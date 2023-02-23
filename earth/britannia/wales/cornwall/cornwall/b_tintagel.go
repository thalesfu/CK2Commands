package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷塔杰尔TintagelBarony struct {
	feud.BaseBarony
}

var BTintagel廷塔杰尔 feud.Barony = &廷塔杰尔TintagelBarony{}

func init() {
	f := BTintagel廷塔杰尔.(*廷塔杰尔TintagelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tintagel",
		TitleName: "廷塔杰尔",
		TitleCode: "b_tintagel",
	}
}
