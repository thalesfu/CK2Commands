package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图韦比亚Et_tuebiaBarony struct {
	feud.BaseBarony
}

var BEt_tuebia图韦比亚 feud.Barony = &图韦比亚Et_tuebiaBarony{}

func init() {
    f := BEt_tuebia图韦比亚.(*图韦比亚Et_tuebiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "et_tuebia",
		TitleName: "图韦比亚",
		TitleCode: "b_et_tuebia",
	}
}
