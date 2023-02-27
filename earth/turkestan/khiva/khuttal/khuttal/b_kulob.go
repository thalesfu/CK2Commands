package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库洛布KulobBarony struct {
	feud.BaseBarony
}

var BKulob库洛布 feud.Barony = &库洛布KulobBarony{}

func init() {
    f := BKulob库洛布.(*库洛布KulobBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kulob",
		TitleName: "库洛布",
		TitleCode: "b_kulob",
	}
}
