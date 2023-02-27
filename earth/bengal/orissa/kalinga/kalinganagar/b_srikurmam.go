package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利俱利蒙SrikurmamBarony struct {
	feud.BaseBarony
}

var BSrikurmam室利俱利蒙 feud.Barony = &室利俱利蒙SrikurmamBarony{}

func init() {
    f := BSrikurmam室利俱利蒙.(*室利俱利蒙SrikurmamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srikurmam",
		TitleName: "室利俱利蒙",
		TitleCode: "b_srikurmam",
	}
}
