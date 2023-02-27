package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶马希TamasiBarony struct {
	feud.BaseBarony
}

var BTamasi陶马希 feud.Barony = &陶马希TamasiBarony{}

func init() {
    f := BTamasi陶马希.(*陶马希TamasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamasi",
		TitleName: "陶马希",
		TitleCode: "b_tamasi",
	}
}
