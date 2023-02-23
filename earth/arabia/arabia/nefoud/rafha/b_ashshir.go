package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔AshshirBarony struct {
	feud.BaseBarony
}

var BAshshir希尔 feud.Barony = &希尔AshshirBarony{}

func init() {
	f := BAshshir希尔.(*希尔AshshirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashshir",
		TitleName: "希尔",
		TitleCode: "b_ashshir",
	}
}
