package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯丝SisBarony struct {
	feud.BaseBarony
}

var BSis斯丝 feud.Barony = &斯丝SisBarony{}

func init() {
    f := BSis斯丝.(*斯丝SisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sis",
		TitleName: "斯丝",
		TitleCode: "b_sis",
	}
}
