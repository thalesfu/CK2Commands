package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼图比MantyubeBarony struct {
	feud.BaseBarony
}

var BMantyube曼图比 feud.Barony = &曼图比MantyubeBarony{}

func init() {
    f := BMantyube曼图比.(*曼图比MantyubeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mantyube",
		TitleName: "曼图比",
		TitleCode: "b_mantyube",
	}
}
