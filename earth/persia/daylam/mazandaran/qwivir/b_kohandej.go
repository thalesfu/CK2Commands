package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科汉德KohandejBarony struct {
	feud.BaseBarony
}

var BKohandej科汉德 feud.Barony = &科汉德KohandejBarony{}

func init() {
    f := BKohandej科汉德.(*科汉德KohandejBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kohandej",
		TitleName: "科汉德",
		TitleCode: "b_kohandej",
	}
}
