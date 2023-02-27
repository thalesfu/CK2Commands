package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兰多AlamdoBarony struct {
	feud.BaseBarony
}

var BAlamdo阿兰多 feud.Barony = &阿兰多AlamdoBarony{}

func init() {
    f := BAlamdo阿兰多.(*阿兰多AlamdoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alamdo",
		TitleName: "阿兰多",
		TitleCode: "b_alamdo",
	}
}
