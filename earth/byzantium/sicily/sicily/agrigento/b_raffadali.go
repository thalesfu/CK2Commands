package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉法达利RaffadaliBarony struct {
	feud.BaseBarony
}

var BRaffadali拉法达利 feud.Barony = &拉法达利RaffadaliBarony{}

func init() {
    f := BRaffadali拉法达利.(*拉法达利RaffadaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raffadali",
		TitleName: "拉法达利",
		TitleCode: "b_raffadali",
	}
}
