package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔什特普TashtypBarony struct {
	feud.BaseBarony
}

var BTashtyp塔什特普 feud.Barony = &塔什特普TashtypBarony{}

func init() {
    f := BTashtyp塔什特普.(*塔什特普TashtypBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tashtyp",
		TitleName: "塔什特普",
		TitleCode: "b_tashtyp",
	}
}
