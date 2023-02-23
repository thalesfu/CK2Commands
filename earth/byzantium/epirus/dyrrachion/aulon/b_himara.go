package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希马拉HimaraBarony struct {
	feud.BaseBarony
}

var BHimara希马拉 feud.Barony = &希马拉HimaraBarony{}

func init() {
	f := BHimara希马拉.(*希马拉HimaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "himara",
		TitleName: "希马拉",
		TitleCode: "b_himara",
	}
}
