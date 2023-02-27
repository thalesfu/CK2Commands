package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗劳恩费尔德FrauenfeldBarony struct {
	feud.BaseBarony
}

var BFrauenfeld弗劳恩费尔德 feud.Barony = &弗劳恩费尔德FrauenfeldBarony{}

func init() {
    f := BFrauenfeld弗劳恩费尔德.(*弗劳恩费尔德FrauenfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frauenfeld",
		TitleName: "弗劳恩费尔德",
		TitleCode: "b_frauenfeld",
	}
}
