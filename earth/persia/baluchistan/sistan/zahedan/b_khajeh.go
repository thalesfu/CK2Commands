package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克哈吉KhajehBarony struct {
	feud.BaseBarony
}

var BKhajeh克哈吉 feud.Barony = &克哈吉KhajehBarony{}

func init() {
	f := BKhajeh克哈吉.(*克哈吉KhajehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khajeh",
		TitleName: "克哈吉",
		TitleCode: "b_khajeh",
	}
}
