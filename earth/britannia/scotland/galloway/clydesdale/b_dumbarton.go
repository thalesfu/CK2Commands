package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓巴顿DumbartonBarony struct {
	feud.BaseBarony
}

var BDumbarton邓巴顿 feud.Barony = &邓巴顿DumbartonBarony{}

func init() {
    f := BDumbarton邓巴顿.(*邓巴顿DumbartonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dumbarton",
		TitleName: "邓巴顿",
		TitleCode: "b_dumbarton",
	}
}
