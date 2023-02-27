package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马达因MadainBarony struct {
	feud.BaseBarony
}

var BMadain马达因 feud.Barony = &马达因MadainBarony{}

func init() {
    f := BMadain马达因.(*马达因MadainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madain",
		TitleName: "马达因",
		TitleCode: "b_madain",
	}
}
