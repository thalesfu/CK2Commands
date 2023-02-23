package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰尔夫兰科CastelfrancoBarony struct {
	feud.BaseBarony
}

var BCastelfranco卡斯泰尔夫兰科 feud.Barony = &卡斯泰尔夫兰科CastelfrancoBarony{}

func init() {
	f := BCastelfranco卡斯泰尔夫兰科.(*卡斯泰尔夫兰科CastelfrancoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelfranco",
		TitleName: "卡斯泰尔夫兰科",
		TitleCode: "b_castelfranco",
	}
}
