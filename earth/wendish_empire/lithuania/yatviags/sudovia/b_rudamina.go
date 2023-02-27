package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁达米纳RudaminaBarony struct {
	feud.BaseBarony
}

var BRudamina鲁达米纳 feud.Barony = &鲁达米纳RudaminaBarony{}

func init() {
    f := BRudamina鲁达米纳.(*鲁达米纳RudaminaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rudamina",
		TitleName: "鲁达米纳",
		TitleCode: "b_rudamina",
	}
}
