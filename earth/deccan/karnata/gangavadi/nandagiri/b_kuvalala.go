package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠婆剌罗KuvalalaBarony struct {
	feud.BaseBarony
}

var BKuvalala鸠婆剌罗 feud.Barony = &鸠婆剌罗KuvalalaBarony{}

func init() {
	f := BKuvalala鸠婆剌罗.(*鸠婆剌罗KuvalalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuvalala",
		TitleName: "鸠婆剌罗",
		TitleCode: "b_kuvalala",
	}
}
