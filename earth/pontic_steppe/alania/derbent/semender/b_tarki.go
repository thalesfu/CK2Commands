package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔基TarkiBarony struct {
	feud.BaseBarony
}

var BTarki塔尔基 feud.Barony = &塔尔基TarkiBarony{}

func init() {
    f := BTarki塔尔基.(*塔尔基TarkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarki",
		TitleName: "塔尔基",
		TitleCode: "b_tarki",
	}
}
