package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯米拉SmilaBarony struct {
	feud.BaseBarony
}

var BSmila斯米拉 feud.Barony = &斯米拉SmilaBarony{}

func init() {
    f := BSmila斯米拉.(*斯米拉SmilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smila",
		TitleName: "斯米拉",
		TitleCode: "b_smila",
	}
}
