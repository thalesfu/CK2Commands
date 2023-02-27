package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希利KhilBarony struct {
	feud.BaseBarony
}

var BKhil希利 feud.Barony = &希利KhilBarony{}

func init() {
    f := BKhil希利.(*希利KhilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khil",
		TitleName: "希利",
		TitleCode: "b_khil",
	}
}
