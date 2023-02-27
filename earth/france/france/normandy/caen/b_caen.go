package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡昂CaenBarony struct {
	feud.BaseBarony
}

var BCaen卡昂 feud.Barony = &卡昂CaenBarony{}

func init() {
    f := BCaen卡昂.(*卡昂CaenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caen",
		TitleName: "卡昂",
		TitleCode: "b_caen",
	}
}
