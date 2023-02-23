package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧马勒AumaleBarony struct {
	feud.BaseBarony
}

var BAumale欧马勒 feud.Barony = &欧马勒AumaleBarony{}

func init() {
	f := BAumale欧马勒.(*欧马勒AumaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aumale",
		TitleName: "欧马勒",
		TitleCode: "b_aumale",
	}
}
