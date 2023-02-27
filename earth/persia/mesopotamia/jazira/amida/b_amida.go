package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿米达AmidaBarony struct {
	feud.BaseBarony
}

var BAmida阿米达 feud.Barony = &阿米达AmidaBarony{}

func init() {
    f := BAmida阿米达.(*阿米达AmidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amida",
		TitleName: "阿米达",
		TitleCode: "b_amida",
	}
}
