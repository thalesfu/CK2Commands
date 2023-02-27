package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮奥扎PyozaBarony struct {
	feud.BaseBarony
}

var BPyoza皮奥扎 feud.Barony = &皮奥扎PyozaBarony{}

func init() {
    f := BPyoza皮奥扎.(*皮奥扎PyozaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyoza",
		TitleName: "皮奥扎",
		TitleCode: "b_pyoza",
	}
}
