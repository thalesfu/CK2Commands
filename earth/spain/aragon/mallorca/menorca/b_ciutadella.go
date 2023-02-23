package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 休塔德利亚CiutadellaBarony struct {
	feud.BaseBarony
}

var BCiutadella休塔德利亚 feud.Barony = &休塔德利亚CiutadellaBarony{}

func init() {
	f := BCiutadella休塔德利亚.(*休塔德利亚CiutadellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciutadella",
		TitleName: "休塔德利亚",
		TitleCode: "b_ciutadella",
	}
}
