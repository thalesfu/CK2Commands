package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 固勒扎KuljaBarony struct {
	feud.BaseBarony
}

var BKulja固勒扎 feud.Barony = &固勒扎KuljaBarony{}

func init() {
	f := BKulja固勒扎.(*固勒扎KuljaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kulja",
		TitleName: "固勒扎",
		TitleCode: "b_kulja",
	}
}
