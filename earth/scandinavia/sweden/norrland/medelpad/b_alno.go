package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔讷AlnoBarony struct {
	feud.BaseBarony
}

var BAlno阿尔讷 feud.Barony = &阿尔讷AlnoBarony{}

func init() {
    f := BAlno阿尔讷.(*阿尔讷AlnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alno",
		TitleName: "阿尔讷",
		TitleCode: "b_alno",
	}
}
