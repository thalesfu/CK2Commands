package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨盖兹SaqqezBarony struct {
	feud.BaseBarony
}

var BSaqqez萨盖兹 feud.Barony = &萨盖兹SaqqezBarony{}

func init() {
    f := BSaqqez萨盖兹.(*萨盖兹SaqqezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saqqez",
		TitleName: "萨盖兹",
		TitleCode: "b_saqqez",
	}
}
