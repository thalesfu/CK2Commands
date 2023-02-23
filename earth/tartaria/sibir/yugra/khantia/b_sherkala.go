package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍尔卡拉SherkalaBarony struct {
	feud.BaseBarony
}

var BSherkala舍尔卡拉 feud.Barony = &舍尔卡拉SherkalaBarony{}

func init() {
	f := BSherkala舍尔卡拉.(*舍尔卡拉SherkalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sherkala",
		TitleName: "舍尔卡拉",
		TitleCode: "b_sherkala",
	}
}
