package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海拜尔KhaybarBarony struct {
	feud.BaseBarony
}

var BKhaybar海拜尔 feud.Barony = &海拜尔KhaybarBarony{}

func init() {
	f := BKhaybar海拜尔.(*海拜尔KhaybarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaybar",
		TitleName: "海拜尔",
		TitleCode: "b_khaybar",
	}
}
