package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴克尔BakearBarony struct {
	feud.BaseBarony
}

var BBakear巴克尔 feud.Barony = &巴克尔BakearBarony{}

func init() {
	f := BBakear巴克尔.(*巴克尔BakearBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakear",
		TitleName: "巴克尔",
		TitleCode: "b_bakear",
	}
}
