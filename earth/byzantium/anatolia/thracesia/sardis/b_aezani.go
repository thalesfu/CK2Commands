package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾赞尼AezaniBarony struct {
	feud.BaseBarony
}

var BAezani艾赞尼 feud.Barony = &艾赞尼AezaniBarony{}

func init() {
	f := BAezani艾赞尼.(*艾赞尼AezaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aezani",
		TitleName: "艾赞尼",
		TitleCode: "b_aezani",
	}
}
