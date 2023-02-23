package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇拉卡赫QalaikahBarony struct {
	feud.BaseBarony
}

var BQalaikah奇拉卡赫 feud.Barony = &奇拉卡赫QalaikahBarony{}

func init() {
	f := BQalaikah奇拉卡赫.(*奇拉卡赫QalaikahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qalaikah",
		TitleName: "奇拉卡赫",
		TitleCode: "b_qalaikah",
	}
}
