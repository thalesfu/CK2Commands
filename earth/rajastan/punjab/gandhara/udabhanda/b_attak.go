package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塔克AttakBarony struct {
	feud.BaseBarony
}

var BAttak阿塔克 feud.Barony = &阿塔克AttakBarony{}

func init() {
	f := BAttak阿塔克.(*阿塔克AttakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attak",
		TitleName: "阿塔克",
		TitleCode: "b_attak",
	}
}
