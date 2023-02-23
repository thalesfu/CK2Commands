package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马尔胡斯HamarhusBarony struct {
	feud.BaseBarony
}

var BHamarhus哈马尔胡斯 feud.Barony = &哈马尔胡斯HamarhusBarony{}

func init() {
	f := BHamarhus哈马尔胡斯.(*哈马尔胡斯HamarhusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamarhus",
		TitleName: "哈马尔胡斯",
		TitleCode: "b_hamarhus",
	}
}
