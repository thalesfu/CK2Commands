package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海拉姆KhaylamBarony struct {
	feud.BaseBarony
}

var BKhaylam海拉姆 feud.Barony = &海拉姆KhaylamBarony{}

func init() {
    f := BKhaylam海拉姆.(*海拉姆KhaylamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaylam",
		TitleName: "海拉姆",
		TitleCode: "b_khaylam",
	}
}
