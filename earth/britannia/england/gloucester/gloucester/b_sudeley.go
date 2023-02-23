package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 休德利SudeleyBarony struct {
	feud.BaseBarony
}

var BSudeley休德利 feud.Barony = &休德利SudeleyBarony{}

func init() {
	f := BSudeley休德利.(*休德利SudeleyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sudeley",
		TitleName: "休德利",
		TitleCode: "b_sudeley",
	}
}
