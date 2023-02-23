package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门陀罗MandlaBarony struct {
	feud.BaseBarony
}

var BMandla门陀罗 feud.Barony = &门陀罗MandlaBarony{}

func init() {
	f := BMandla门陀罗.(*门陀罗MandlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandla",
		TitleName: "门陀罗",
		TitleCode: "b_mandla",
	}
}
