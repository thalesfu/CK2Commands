package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃戈VagaBarony struct {
	feud.BaseBarony
}

var BVaga沃戈 feud.Barony = &沃戈VagaBarony{}

func init() {
	f := BVaga沃戈.(*沃戈VagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaga",
		TitleName: "沃戈",
		TitleCode: "b_vaga",
	}
}
