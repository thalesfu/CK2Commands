package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 屈恩灵KuenringBarony struct {
	feud.BaseBarony
}

var BKuenring屈恩灵 feud.Barony = &屈恩灵KuenringBarony{}

func init() {
	f := BKuenring屈恩灵.(*屈恩灵KuenringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuenring",
		TitleName: "屈恩灵",
		TitleCode: "b_kuenring",
	}
}
