package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫吉奥MelgueilBarony struct {
	feud.BaseBarony
}

var BMelgueil莫吉奥 feud.Barony = &莫吉奥MelgueilBarony{}

func init() {
	f := BMelgueil莫吉奥.(*莫吉奥MelgueilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melgueil",
		TitleName: "莫吉奥",
		TitleCode: "b_melgueil",
	}
}
