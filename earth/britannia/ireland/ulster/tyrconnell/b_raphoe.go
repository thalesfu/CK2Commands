package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉福RaphoeBarony struct {
	feud.BaseBarony
}

var BRaphoe拉福 feud.Barony = &拉福RaphoeBarony{}

func init() {
	f := BRaphoe拉福.(*拉福RaphoeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raphoe",
		TitleName: "拉福",
		TitleCode: "b_raphoe",
	}
}
