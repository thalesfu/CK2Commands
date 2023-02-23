package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍伊KhoyBarony struct {
	feud.BaseBarony
}

var BKhoy霍伊 feud.Barony = &霍伊KhoyBarony{}

func init() {
	f := BKhoy霍伊.(*霍伊KhoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khoy",
		TitleName: "霍伊",
		TitleCode: "b_khoy",
	}
}
