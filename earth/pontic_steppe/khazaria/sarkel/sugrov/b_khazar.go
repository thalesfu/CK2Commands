package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 可萨KhazarBarony struct {
	feud.BaseBarony
}

var BKhazar可萨 feud.Barony = &可萨KhazarBarony{}

func init() {
    f := BKhazar可萨.(*可萨KhazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khazar",
		TitleName: "可萨",
		TitleCode: "b_khazar",
	}
}
