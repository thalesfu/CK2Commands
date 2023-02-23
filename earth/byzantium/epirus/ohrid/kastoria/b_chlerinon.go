package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克勒里农ChlerinonBarony struct {
	feud.BaseBarony
}

var BChlerinon克勒里农 feud.Barony = &克勒里农ChlerinonBarony{}

func init() {
	f := BChlerinon克勒里农.(*克勒里农ChlerinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chlerinon",
		TitleName: "克勒里农",
		TitleCode: "b_chlerinon",
	}
}
