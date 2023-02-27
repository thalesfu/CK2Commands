package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱乌卡LeucaBarony struct {
	feud.BaseBarony
}

var BLeuca莱乌卡 feud.Barony = &莱乌卡LeucaBarony{}

func init() {
    f := BLeuca莱乌卡.(*莱乌卡LeucaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leuca",
		TitleName: "莱乌卡",
		TitleCode: "b_leuca",
	}
}
