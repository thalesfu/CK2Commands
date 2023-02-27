package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉贝特南吉LapytnangkBarony struct {
	feud.BaseBarony
}

var BLapytnangk拉贝特南吉 feud.Barony = &拉贝特南吉LapytnangkBarony{}

func init() {
    f := BLapytnangk拉贝特南吉.(*拉贝特南吉LapytnangkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lapytnangk",
		TitleName: "拉贝特南吉",
		TitleCode: "b_lapytnangk",
	}
}
