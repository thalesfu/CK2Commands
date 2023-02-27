package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格德AgdeBarony struct {
	feud.BaseBarony
}

var BAgde阿格德 feud.Barony = &阿格德AgdeBarony{}

func init() {
    f := BAgde阿格德.(*阿格德AgdeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agde",
		TitleName: "阿格德",
		TitleCode: "b_agde",
	}
}
