package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希瓦KhivaBarony struct {
	feud.BaseBarony
}

var BKhiva希瓦 feud.Barony = &希瓦KhivaBarony{}

func init() {
	f := BKhiva希瓦.(*希瓦KhivaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khiva",
		TitleName: "希瓦",
		TitleCode: "b_khiva",
	}
}
