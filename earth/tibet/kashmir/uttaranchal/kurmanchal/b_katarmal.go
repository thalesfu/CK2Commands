package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦吒罗末罗KatarmalBarony struct {
	feud.BaseBarony
}

var BKatarmal迦吒罗末罗 feud.Barony = &迦吒罗末罗KatarmalBarony{}

func init() {
	f := BKatarmal迦吒罗末罗.(*迦吒罗末罗KatarmalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katarmal",
		TitleName: "迦吒罗末罗",
		TitleCode: "b_katarmal",
	}
}
