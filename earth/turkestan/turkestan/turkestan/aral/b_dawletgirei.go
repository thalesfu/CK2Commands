package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达夫列特_吉雷DawletgireiBarony struct {
	feud.BaseBarony
}

var BDawletgirei达夫列特_吉雷 feud.Barony = &达夫列特_吉雷DawletgireiBarony{}

func init() {
    f := BDawletgirei达夫列特_吉雷.(*达夫列特_吉雷DawletgireiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dawletgirei",
		TitleName: "达夫列特_吉雷",
		TitleCode: "b_dawletgirei",
	}
}
