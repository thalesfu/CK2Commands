package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔卡吕德MarkarydBarony struct {
	feud.BaseBarony
}

var BMarkaryd马尔卡吕德 feud.Barony = &马尔卡吕德MarkarydBarony{}

func init() {
	f := BMarkaryd马尔卡吕德.(*马尔卡吕德MarkarydBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "markaryd",
		TitleName: "马尔卡吕德",
		TitleCode: "b_markaryd",
	}
}
