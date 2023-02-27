package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里古达尔兹AligoodarzBarony struct {
	feud.BaseBarony
}

var BAligoodarz阿里古达尔兹 feud.Barony = &阿里古达尔兹AligoodarzBarony{}

func init() {
    f := BAligoodarz阿里古达尔兹.(*阿里古达尔兹AligoodarzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aligoodarz",
		TitleName: "阿里古达尔兹",
		TitleCode: "b_aligoodarz",
	}
}
