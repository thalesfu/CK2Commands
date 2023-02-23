package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾赫罗姆JahromBarony struct {
	feud.BaseBarony
}

var BJahrom贾赫罗姆 feud.Barony = &贾赫罗姆JahromBarony{}

func init() {
	f := BJahrom贾赫罗姆.(*贾赫罗姆JahromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jahrom",
		TitleName: "贾赫罗姆",
		TitleCode: "b_jahrom",
	}
}
