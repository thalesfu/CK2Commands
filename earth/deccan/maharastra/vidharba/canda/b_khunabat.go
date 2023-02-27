package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘那跋KhunabatBarony struct {
	feud.BaseBarony
}

var BKhunabat丘那跋 feud.Barony = &丘那跋KhunabatBarony{}

func init() {
    f := BKhunabat丘那跋.(*丘那跋KhunabatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khunabat",
		TitleName: "丘那跋",
		TitleCode: "b_khunabat",
	}
}
