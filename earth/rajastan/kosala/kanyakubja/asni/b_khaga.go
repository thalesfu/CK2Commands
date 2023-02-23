package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡迦KhagaBarony struct {
	feud.BaseBarony
}

var BKhaga卡迦 feud.Barony = &卡迦KhagaBarony{}

func init() {
	f := BKhaga卡迦.(*卡迦KhagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaga",
		TitleName: "卡迦",
		TitleCode: "b_khaga",
	}
}
