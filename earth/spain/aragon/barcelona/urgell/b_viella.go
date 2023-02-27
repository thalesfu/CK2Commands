package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别利亚ViellaBarony struct {
	feud.BaseBarony
}

var BViella别利亚 feud.Barony = &别利亚ViellaBarony{}

func init() {
    f := BViella别利亚.(*别利亚ViellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viella",
		TitleName: "别利亚",
		TitleCode: "b_viella",
	}
}
