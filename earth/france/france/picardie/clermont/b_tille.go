package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂耶TilleBarony struct {
	feud.BaseBarony
}

var BTille蒂耶 feud.Barony = &蒂耶TilleBarony{}

func init() {
    f := BTille蒂耶.(*蒂耶TilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tille",
		TitleName: "蒂耶",
		TitleCode: "b_tille",
	}
}
