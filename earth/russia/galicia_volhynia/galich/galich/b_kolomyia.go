package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科洛梅亚KolomyiaBarony struct {
	feud.BaseBarony
}

var BKolomyia科洛梅亚 feud.Barony = &科洛梅亚KolomyiaBarony{}

func init() {
    f := BKolomyia科洛梅亚.(*科洛梅亚KolomyiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolomyia",
		TitleName: "科洛梅亚",
		TitleCode: "b_kolomyia",
	}
}
