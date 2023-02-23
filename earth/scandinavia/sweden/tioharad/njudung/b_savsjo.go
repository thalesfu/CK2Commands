package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛夫舍SavsjoBarony struct {
	feud.BaseBarony
}

var BSavsjo赛夫舍 feud.Barony = &赛夫舍SavsjoBarony{}

func init() {
	f := BSavsjo赛夫舍.(*赛夫舍SavsjoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savsjo",
		TitleName: "赛夫舍",
		TitleCode: "b_savsjo",
	}
}
