package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔韦泰TerveteBarony struct {
	feud.BaseBarony
}

var BTervete泰尔韦泰 feud.Barony = &泰尔韦泰TerveteBarony{}

func init() {
    f := BTervete泰尔韦泰.(*泰尔韦泰TerveteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tervete",
		TitleName: "泰尔韦泰",
		TitleCode: "b_tervete",
	}
}
