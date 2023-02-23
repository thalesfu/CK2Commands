package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卢佐SaluzzoBarony struct {
	feud.BaseBarony
}

var BSaluzzo萨卢佐 feud.Barony = &萨卢佐SaluzzoBarony{}

func init() {
	f := BSaluzzo萨卢佐.(*萨卢佐SaluzzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saluzzo",
		TitleName: "萨卢佐",
		TitleCode: "b_saluzzo",
	}
}
