package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图斯TusBarony struct {
	feud.BaseBarony
}

var BTus图斯 feud.Barony = &图斯TusBarony{}

func init() {
    f := BTus图斯.(*图斯TusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tus",
		TitleName: "图斯",
		TitleCode: "b_tus",
	}
}
