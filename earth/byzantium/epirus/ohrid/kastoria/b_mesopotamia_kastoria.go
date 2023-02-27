package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 美索不达米亚Mesopotamia_kastoriaBarony struct {
	feud.BaseBarony
}

var BMesopotamia_kastoria美索不达米亚 feud.Barony = &美索不达米亚Mesopotamia_kastoriaBarony{}

func init() {
    f := BMesopotamia_kastoria美索不达米亚.(*美索不达米亚Mesopotamia_kastoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mesopotamia_kastoria",
		TitleName: "美索不达米亚",
		TitleCode: "b_mesopotamia_kastoria",
	}
}
