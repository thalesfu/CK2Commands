package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索林SolinBarony struct {
	feud.BaseBarony
}

var BSolin索林 feud.Barony = &索林SolinBarony{}

func init() {
    f := BSolin索林.(*索林SolinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solin",
		TitleName: "索林",
		TitleCode: "b_solin",
	}
}
