package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉费GalafiBarony struct {
	feud.BaseBarony
}

var BGalafi加拉费 feud.Barony = &加拉费GalafiBarony{}

func init() {
    f := BGalafi加拉费.(*加拉费GalafiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galafi",
		TitleName: "加拉费",
		TitleCode: "b_galafi",
	}
}
