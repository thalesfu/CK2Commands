package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔卡维什基斯VilkaviskisBarony struct {
	feud.BaseBarony
}

var BVilkaviskis维尔卡维什基斯 feud.Barony = &维尔卡维什基斯VilkaviskisBarony{}

func init() {
    f := BVilkaviskis维尔卡维什基斯.(*维尔卡维什基斯VilkaviskisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilkaviskis",
		TitleName: "维尔卡维什基斯",
		TitleCode: "b_vilkaviskis",
	}
}
