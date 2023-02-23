package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰迪南LlandinamBarony struct {
	feud.BaseBarony
}

var BLlandinam兰迪南 feud.Barony = &兰迪南LlandinamBarony{}

func init() {
	f := BLlandinam兰迪南.(*兰迪南LlandinamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llandinam",
		TitleName: "兰迪南",
		TitleCode: "b_llandinam",
	}
}
