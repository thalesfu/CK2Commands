package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼荼丽湿伐罗MandleshwarBarony struct {
	feud.BaseBarony
}

var BMandleshwar曼荼丽湿伐罗 feud.Barony = &曼荼丽湿伐罗MandleshwarBarony{}

func init() {
	f := BMandleshwar曼荼丽湿伐罗.(*曼荼丽湿伐罗MandleshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandleshwar",
		TitleName: "曼荼丽湿伐罗",
		TitleCode: "b_mandleshwar",
	}
}
