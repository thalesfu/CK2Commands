package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泥弥沙NaimisaBarony struct {
	feud.BaseBarony
}

var BNaimisa泥弥沙 feud.Barony = &泥弥沙NaimisaBarony{}

func init() {
	f := BNaimisa泥弥沙.(*泥弥沙NaimisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naimisa",
		TitleName: "泥弥沙",
		TitleCode: "b_naimisa",
	}
}
