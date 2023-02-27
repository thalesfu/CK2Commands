package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科扬德KoyandyBarony struct {
	feud.BaseBarony
}

var BKoyandy科扬德 feud.Barony = &科扬德KoyandyBarony{}

func init() {
    f := BKoyandy科扬德.(*科扬德KoyandyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koyandy",
		TitleName: "科扬德",
		TitleCode: "b_koyandy",
	}
}
