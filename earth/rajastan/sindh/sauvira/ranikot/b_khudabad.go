package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘陀跋KhudabadBarony struct {
	feud.BaseBarony
}

var BKhudabad丘陀跋 feud.Barony = &丘陀跋KhudabadBarony{}

func init() {
    f := BKhudabad丘陀跋.(*丘陀跋KhudabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khudabad",
		TitleName: "丘陀跋",
		TitleCode: "b_khudabad",
	}
}
