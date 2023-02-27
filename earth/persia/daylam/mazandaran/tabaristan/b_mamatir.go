package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛马提尔MamatirBarony struct {
	feud.BaseBarony
}

var BMamatir玛马提尔 feud.Barony = &玛马提尔MamatirBarony{}

func init() {
    f := BMamatir玛马提尔.(*玛马提尔MamatirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamatir",
		TitleName: "玛马提尔",
		TitleCode: "b_mamatir",
	}
}
