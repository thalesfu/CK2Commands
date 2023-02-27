package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克_梅切特AkmechetBarony struct {
	feud.BaseBarony
}

var BAkmechet阿克_梅切特 feud.Barony = &阿克_梅切特AkmechetBarony{}

func init() {
    f := BAkmechet阿克_梅切特.(*阿克_梅切特AkmechetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akmechet",
		TitleName: "阿克_梅切特",
		TitleCode: "b_akmechet",
	}
}
