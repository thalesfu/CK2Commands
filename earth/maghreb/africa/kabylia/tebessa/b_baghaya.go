package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴加耶BaghayaBarony struct {
	feud.BaseBarony
}

var BBaghaya巴加耶 feud.Barony = &巴加耶BaghayaBarony{}

func init() {
    f := BBaghaya巴加耶.(*巴加耶BaghayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghaya",
		TitleName: "巴加耶",
		TitleCode: "b_baghaya",
	}
}
