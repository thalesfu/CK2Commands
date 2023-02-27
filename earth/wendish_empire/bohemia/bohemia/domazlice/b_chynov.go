package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希诺夫ChynovBarony struct {
	feud.BaseBarony
}

var BChynov希诺夫 feud.Barony = &希诺夫ChynovBarony{}

func init() {
    f := BChynov希诺夫.(*希诺夫ChynovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chynov",
		TitleName: "希诺夫",
		TitleCode: "b_chynov",
	}
}
