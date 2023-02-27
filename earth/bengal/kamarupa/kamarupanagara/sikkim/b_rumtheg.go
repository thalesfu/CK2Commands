package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆德RumthegBarony struct {
	feud.BaseBarony
}

var BRumtheg隆德 feud.Barony = &隆德RumthegBarony{}

func init() {
    f := BRumtheg隆德.(*隆德RumthegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rumtheg",
		TitleName: "隆德",
		TitleCode: "b_rumtheg",
	}
}
