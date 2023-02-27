package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比谢克特BesiktyBarony struct {
	feud.BaseBarony
}

var BBesikty比谢克特 feud.Barony = &比谢克特BesiktyBarony{}

func init() {
    f := BBesikty比谢克特.(*比谢克特BesiktyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besikty",
		TitleName: "比谢克特",
		TitleCode: "b_besikty",
	}
}
