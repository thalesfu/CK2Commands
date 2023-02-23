package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库姆塔格KumtagBarony struct {
	feud.BaseBarony
}

var BKumtag库姆塔格 feud.Barony = &库姆塔格KumtagBarony{}

func init() {
	f := BKumtag库姆塔格.(*库姆塔格KumtagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumtag",
		TitleName: "库姆塔格",
		TitleCode: "b_kumtag",
	}
}
