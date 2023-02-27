package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺贡迪Nok_kundiBarony struct {
	feud.BaseBarony
}

var BNok_kundi诺贡迪 feud.Barony = &诺贡迪Nok_kundiBarony{}

func init() {
    f := BNok_kundi诺贡迪.(*诺贡迪Nok_kundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nok_kundi",
		TitleName: "诺贡迪",
		TitleCode: "b_nok_kundi",
	}
}
