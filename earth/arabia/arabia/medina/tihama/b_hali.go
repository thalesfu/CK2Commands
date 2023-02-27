package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈利HaliBarony struct {
	feud.BaseBarony
}

var BHali哈利 feud.Barony = &哈利HaliBarony{}

func init() {
    f := BHali哈利.(*哈利HaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hali",
		TitleName: "哈利",
		TitleCode: "b_hali",
	}
}
