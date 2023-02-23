package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肖蒙ChaumontBarony struct {
	feud.BaseBarony
}

var BChaumont肖蒙 feud.Barony = &肖蒙ChaumontBarony{}

func init() {
	f := BChaumont肖蒙.(*肖蒙ChaumontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaumont",
		TitleName: "肖蒙",
		TitleCode: "b_chaumont",
	}
}
