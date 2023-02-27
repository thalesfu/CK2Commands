package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨马尔SamarBarony struct {
	feud.BaseBarony
}

var BSamar萨马尔 feud.Barony = &萨马尔SamarBarony{}

func init() {
    f := BSamar萨马尔.(*萨马尔SamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samar",
		TitleName: "萨马尔",
		TitleCode: "b_samar",
	}
}
