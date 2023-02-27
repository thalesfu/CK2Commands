package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞琉西亚SeleukeiaBarony struct {
	feud.BaseBarony
}

var BSeleukeia塞琉西亚 feud.Barony = &塞琉西亚SeleukeiaBarony{}

func init() {
    f := BSeleukeia塞琉西亚.(*塞琉西亚SeleukeiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seleukeia",
		TitleName: "塞琉西亚",
		TitleCode: "b_seleukeia",
	}
}
