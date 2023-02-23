package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡朵KhongtogBarony struct {
	feud.BaseBarony
}

var BKhongtog贡朵 feud.Barony = &贡朵KhongtogBarony{}

func init() {
	f := BKhongtog贡朵.(*贡朵KhongtogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khongtog",
		TitleName: "贡朵",
		TitleCode: "b_khongtog",
	}
}
