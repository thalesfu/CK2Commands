package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷茨RetzBarony struct {
	feud.BaseBarony
}

var BRetz雷茨 feud.Barony = &雷茨RetzBarony{}

func init() {
	f := BRetz雷茨.(*雷茨RetzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "retz",
		TitleName: "雷茨",
		TitleCode: "b_retz",
	}
}
