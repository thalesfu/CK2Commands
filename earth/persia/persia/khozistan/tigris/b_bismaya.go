package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯玛亚BismayaBarony struct {
	feud.BaseBarony
}

var BBismaya比斯玛亚 feud.Barony = &比斯玛亚BismayaBarony{}

func init() {
    f := BBismaya比斯玛亚.(*比斯玛亚BismayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bismaya",
		TitleName: "比斯玛亚",
		TitleCode: "b_bismaya",
	}
}
