package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利斯多斯KarystosBarony struct {
	feud.BaseBarony
}

var BKarystos卡利斯多斯 feud.Barony = &卡利斯多斯KarystosBarony{}

func init() {
    f := BKarystos卡利斯多斯.(*卡利斯多斯KarystosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karystos",
		TitleName: "卡利斯多斯",
		TitleCode: "b_karystos",
	}
}
