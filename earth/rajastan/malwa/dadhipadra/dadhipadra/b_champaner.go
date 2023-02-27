package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚庞ChampanerBarony struct {
	feud.BaseBarony
}

var BChampaner尚庞 feud.Barony = &尚庞ChampanerBarony{}

func init() {
    f := BChampaner尚庞.(*尚庞ChampanerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "champaner",
		TitleName: "尚庞",
		TitleCode: "b_champaner",
	}
}
