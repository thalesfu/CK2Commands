package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琉刻LeuceBarony struct {
	feud.BaseBarony
}

var BLeuce琉刻 feud.Barony = &琉刻LeuceBarony{}

func init() {
	f := BLeuce琉刻.(*琉刻LeuceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leuce",
		TitleName: "琉刻",
		TitleCode: "b_leuce",
	}
}
