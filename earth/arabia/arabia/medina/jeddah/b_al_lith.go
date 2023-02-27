package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯Al_lithBarony struct {
	feud.BaseBarony
}

var BAl_lith利斯 feud.Barony = &利斯Al_lithBarony{}

func init() {
    f := BAl_lith利斯.(*利斯Al_lithBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_lith",
		TitleName: "利斯",
		TitleCode: "b_al_lith",
	}
}
