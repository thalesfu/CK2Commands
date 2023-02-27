package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利夫LiwBarony struct {
	feud.BaseBarony
}

var BLiw利夫 feud.Barony = &利夫LiwBarony{}

func init() {
    f := BLiw利夫.(*利夫LiwBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liw",
		TitleName: "利夫",
		TitleCode: "b_liw",
	}
}
