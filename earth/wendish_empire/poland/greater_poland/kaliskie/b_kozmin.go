package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科伊明KozminBarony struct {
	feud.BaseBarony
}

var BKozmin科伊明 feud.Barony = &科伊明KozminBarony{}

func init() {
    f := BKozmin科伊明.(*科伊明KozminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozmin",
		TitleName: "科伊明",
		TitleCode: "b_kozmin",
	}
}
