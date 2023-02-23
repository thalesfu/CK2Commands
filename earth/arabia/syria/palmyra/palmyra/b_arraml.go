package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖姆莱ArramlBarony struct {
	feud.BaseBarony
}

var BArraml赖姆莱 feud.Barony = &赖姆莱ArramlBarony{}

func init() {
	f := BArraml赖姆莱.(*赖姆莱ArramlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arraml",
		TitleName: "赖姆莱",
		TitleCode: "b_arraml",
	}
}
