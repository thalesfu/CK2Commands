package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣索沃尔_昂皮赛StsauveurenpuisayeBarony struct {
	feud.BaseBarony
}

var BStsauveurenpuisaye圣索沃尔_昂皮赛 feud.Barony = &圣索沃尔_昂皮赛StsauveurenpuisayeBarony{}

func init() {
	f := BStsauveurenpuisaye圣索沃尔_昂皮赛.(*圣索沃尔_昂皮赛StsauveurenpuisayeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stsauveurenpuisaye",
		TitleName: "圣索沃尔_昂皮赛",
		TitleCode: "b_stsauveurenpuisaye",
	}
}
