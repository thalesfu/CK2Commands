package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久姆拉JumlaBarony struct {
	feud.BaseBarony
}

var BJumla久姆拉 feud.Barony = &久姆拉JumlaBarony{}

func init() {
	f := BJumla久姆拉.(*久姆拉JumlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jumla",
		TitleName: "久姆拉",
		TitleCode: "b_jumla",
	}
}
