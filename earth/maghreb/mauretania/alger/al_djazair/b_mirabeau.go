package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉本海达MirabeauBarony struct {
	feud.BaseBarony
}

var BMirabeau德拉本海达 feud.Barony = &德拉本海达MirabeauBarony{}

func init() {
    f := BMirabeau德拉本海达.(*德拉本海达MirabeauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirabeau",
		TitleName: "德拉本海达",
		TitleCode: "b_mirabeau",
	}
}
