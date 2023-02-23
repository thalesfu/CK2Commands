package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克季KetBarony struct {
	feud.BaseBarony
}

var BKet克季 feud.Barony = &克季KetBarony{}

func init() {
	f := BKet克季.(*克季KetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ket",
		TitleName: "克季",
		TitleCode: "b_ket",
	}
}
