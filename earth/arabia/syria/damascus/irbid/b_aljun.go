package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿杰隆AljunBarony struct {
	feud.BaseBarony
}

var BAljun阿杰隆 feud.Barony = &阿杰隆AljunBarony{}

func init() {
    f := BAljun阿杰隆.(*阿杰隆AljunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljun",
		TitleName: "阿杰隆",
		TitleCode: "b_aljun",
	}
}
