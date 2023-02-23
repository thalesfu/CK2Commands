package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古城GuchengBarony struct {
	feud.BaseBarony
}

var BGucheng古城 feud.Barony = &古城GuchengBarony{}

func init() {
	f := BGucheng古城.(*古城GuchengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gucheng",
		TitleName: "古城",
		TitleCode: "b_gucheng",
	}
}
