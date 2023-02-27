package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩里塞奥里翁PeritheorionBarony struct {
	feud.BaseBarony
}

var BPeritheorion佩里塞奥里翁 feud.Barony = &佩里塞奥里翁PeritheorionBarony{}

func init() {
    f := BPeritheorion佩里塞奥里翁.(*佩里塞奥里翁PeritheorionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peritheorion",
		TitleName: "佩里塞奥里翁",
		TitleCode: "b_peritheorion",
	}
}
