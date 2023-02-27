package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿韦讷AvesnesBarony struct {
	feud.BaseBarony
}

var BAvesnes阿韦讷 feud.Barony = &阿韦讷AvesnesBarony{}

func init() {
    f := BAvesnes阿韦讷.(*阿韦讷AvesnesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avesnes",
		TitleName: "阿韦讷",
		TitleCode: "b_avesnes",
	}
}
