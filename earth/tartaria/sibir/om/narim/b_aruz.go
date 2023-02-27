package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿鲁兹AruzBarony struct {
	feud.BaseBarony
}

var BAruz阿鲁兹 feud.Barony = &阿鲁兹AruzBarony{}

func init() {
    f := BAruz阿鲁兹.(*阿鲁兹AruzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aruz",
		TitleName: "阿鲁兹",
		TitleCode: "b_aruz",
	}
}
