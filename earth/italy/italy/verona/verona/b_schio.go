package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯基奥SchioBarony struct {
	feud.BaseBarony
}

var BSchio斯基奥 feud.Barony = &斯基奥SchioBarony{}

func init() {
	f := BSchio斯基奥.(*斯基奥SchioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schio",
		TitleName: "斯基奥",
		TitleCode: "b_schio",
	}
}
