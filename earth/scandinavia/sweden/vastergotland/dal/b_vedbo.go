package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦德布VedboBarony struct {
	feud.BaseBarony
}

var BVedbo韦德布 feud.Barony = &韦德布VedboBarony{}

func init() {
    f := BVedbo韦德布.(*韦德布VedboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vedbo",
		TitleName: "韦德布",
		TitleCode: "b_vedbo",
	}
}
