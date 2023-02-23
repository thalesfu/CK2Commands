package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布蒙多BoumoundoBarony struct {
	feud.BaseBarony
}

var BBoumoundo布蒙多 feud.Barony = &布蒙多BoumoundoBarony{}

func init() {
	f := BBoumoundo布蒙多.(*布蒙多BoumoundoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boumoundo",
		TitleName: "布蒙多",
		TitleCode: "b_boumoundo",
	}
}
