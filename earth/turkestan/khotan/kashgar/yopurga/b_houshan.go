package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 后善HoushanBarony struct {
	feud.BaseBarony
}

var BHoushan后善 feud.Barony = &后善HoushanBarony{}

func init() {
	f := BHoushan后善.(*后善HoushanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "houshan",
		TitleName: "后善",
		TitleCode: "b_houshan",
	}
}
