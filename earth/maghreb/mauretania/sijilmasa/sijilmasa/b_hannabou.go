package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉纳布HannabouBarony struct {
	feud.BaseBarony
}

var BHannabou汉纳布 feud.Barony = &汉纳布HannabouBarony{}

func init() {
	f := BHannabou汉纳布.(*汉纳布HannabouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hannabou",
		TitleName: "汉纳布",
		TitleCode: "b_hannabou",
	}
}
