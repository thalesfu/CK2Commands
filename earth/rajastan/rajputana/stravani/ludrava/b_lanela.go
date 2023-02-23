package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗内罗LanelaBarony struct {
	feud.BaseBarony
}

var BLanela罗内罗 feud.Barony = &罗内罗LanelaBarony{}

func init() {
	f := BLanela罗内罗.(*罗内罗LanelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lanela",
		TitleName: "罗内罗",
		TitleCode: "b_lanela",
	}
}
