package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥列夫斯克OlevskBarony struct {
	feud.BaseBarony
}

var BOlevsk奥列夫斯克 feud.Barony = &奥列夫斯克OlevskBarony{}

func init() {
	f := BOlevsk奥列夫斯克.(*奥列夫斯克OlevskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olevsk",
		TitleName: "奥列夫斯克",
		TitleCode: "b_olevsk",
	}
}
