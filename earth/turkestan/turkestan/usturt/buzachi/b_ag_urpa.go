package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格_乌尔帕Ag_urpaBarony struct {
	feud.BaseBarony
}

var BAg_urpa阿格_乌尔帕 feud.Barony = &阿格_乌尔帕Ag_urpaBarony{}

func init() {
    f := BAg_urpa阿格_乌尔帕.(*阿格_乌尔帕Ag_urpaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ag_urpa",
		TitleName: "阿格_乌尔帕",
		TitleCode: "b_ag_urpa",
	}
}
