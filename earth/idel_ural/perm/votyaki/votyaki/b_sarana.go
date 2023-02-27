package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉纳SaranaBarony struct {
	feud.BaseBarony
}

var BSarana萨拉纳 feud.Barony = &萨拉纳SaranaBarony{}

func init() {
    f := BSarana萨拉纳.(*萨拉纳SaranaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarana",
		TitleName: "萨拉纳",
		TitleCode: "b_sarana",
	}
}
