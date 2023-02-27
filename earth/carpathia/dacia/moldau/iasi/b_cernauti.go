package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔讷乌齐CernautiBarony struct {
	feud.BaseBarony
}

var BCernauti切尔讷乌齐 feud.Barony = &切尔讷乌齐CernautiBarony{}

func init() {
    f := BCernauti切尔讷乌齐.(*切尔讷乌齐CernautiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cernauti",
		TitleName: "切尔讷乌齐",
		TitleCode: "b_cernauti",
	}
}
