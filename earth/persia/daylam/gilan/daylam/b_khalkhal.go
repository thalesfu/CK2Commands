package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈勒哈勒KhalkhalBarony struct {
	feud.BaseBarony
}

var BKhalkhal哈勒哈勒 feud.Barony = &哈勒哈勒KhalkhalBarony{}

func init() {
	f := BKhalkhal哈勒哈勒.(*哈勒哈勒KhalkhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khalkhal",
		TitleName: "哈勒哈勒",
		TitleCode: "b_khalkhal",
	}
}
