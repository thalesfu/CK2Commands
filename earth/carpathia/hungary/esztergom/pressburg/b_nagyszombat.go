package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉松鲍特NagyszombatBarony struct {
	feud.BaseBarony
}

var BNagyszombat瑙吉松鲍特 feud.Barony = &瑙吉松鲍特NagyszombatBarony{}

func init() {
    f := BNagyszombat瑙吉松鲍特.(*瑙吉松鲍特NagyszombatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagyszombat",
		TitleName: "瑙吉松鲍特",
		TitleCode: "b_nagyszombat",
	}
}
