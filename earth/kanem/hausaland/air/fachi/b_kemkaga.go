package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康卡加KemkagaBarony struct {
	feud.BaseBarony
}

var BKemkaga康卡加 feud.Barony = &康卡加KemkagaBarony{}

func init() {
    f := BKemkaga康卡加.(*康卡加KemkagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kemkaga",
		TitleName: "康卡加",
		TitleCode: "b_kemkaga",
	}
}
