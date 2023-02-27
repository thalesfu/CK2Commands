package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苦恰KuqaBarony struct {
	feud.BaseBarony
}

var BKuqa苦恰 feud.Barony = &苦恰KuqaBarony{}

func init() {
    f := BKuqa苦恰.(*苦恰KuqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuqa",
		TitleName: "苦恰",
		TitleCode: "b_kuqa",
	}
}
