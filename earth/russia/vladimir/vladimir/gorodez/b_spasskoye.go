package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯帕斯科耶SpasskoyeBarony struct {
	feud.BaseBarony
}

var BSpasskoye斯帕斯科耶 feud.Barony = &斯帕斯科耶SpasskoyeBarony{}

func init() {
    f := BSpasskoye斯帕斯科耶.(*斯帕斯科耶SpasskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spasskoye",
		TitleName: "斯帕斯科耶",
		TitleCode: "b_spasskoye",
	}
}
