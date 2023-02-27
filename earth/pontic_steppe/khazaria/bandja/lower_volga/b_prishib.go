package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里希布PrishibBarony struct {
	feud.BaseBarony
}

var BPrishib普里希布 feud.Barony = &普里希布PrishibBarony{}

func init() {
    f := BPrishib普里希布.(*普里希布PrishibBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prishib",
		TitleName: "普里希布",
		TitleCode: "b_prishib",
	}
}
