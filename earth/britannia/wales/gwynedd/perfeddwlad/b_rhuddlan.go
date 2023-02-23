package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里兹兰RhuddlanBarony struct {
	feud.BaseBarony
}

var BRhuddlan里兹兰 feud.Barony = &里兹兰RhuddlanBarony{}

func init() {
	f := BRhuddlan里兹兰.(*里兹兰RhuddlanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhuddlan",
		TitleName: "里兹兰",
		TitleCode: "b_rhuddlan",
	}
}
