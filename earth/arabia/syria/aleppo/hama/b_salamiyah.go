package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞莱米耶SalamiyahBarony struct {
	feud.BaseBarony
}

var BSalamiyah塞莱米耶 feud.Barony = &塞莱米耶SalamiyahBarony{}

func init() {
    f := BSalamiyah塞莱米耶.(*塞莱米耶SalamiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salamiyah",
		TitleName: "塞莱米耶",
		TitleCode: "b_salamiyah",
	}
}
