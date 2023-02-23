package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马德第MaddertyBarony struct {
	feud.BaseBarony
}

var BMadderty马德第 feud.Barony = &马德第MaddertyBarony{}

func init() {
	f := BMadderty马德第.(*马德第MaddertyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madderty",
		TitleName: "马德第",
		TitleCode: "b_madderty",
	}
}
