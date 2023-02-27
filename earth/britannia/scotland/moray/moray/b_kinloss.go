package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 金洛斯KinlossBarony struct {
	feud.BaseBarony
}

var BKinloss金洛斯 feud.Barony = &金洛斯KinlossBarony{}

func init() {
    f := BKinloss金洛斯.(*金洛斯KinlossBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinloss",
		TitleName: "金洛斯",
		TitleCode: "b_kinloss",
	}
}
