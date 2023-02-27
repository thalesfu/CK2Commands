package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈登堡NeidenburgBarony struct {
	feud.BaseBarony
}

var BNeidenburg奈登堡 feud.Barony = &奈登堡NeidenburgBarony{}

func init() {
    f := BNeidenburg奈登堡.(*奈登堡NeidenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neidenburg",
		TitleName: "奈登堡",
		TitleCode: "b_neidenburg",
	}
}
