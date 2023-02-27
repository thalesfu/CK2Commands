package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗昂RoanneBarony struct {
	feud.BaseBarony
}

var BRoanne罗昂 feud.Barony = &罗昂RoanneBarony{}

func init() {
    f := BRoanne罗昂.(*罗昂RoanneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roanne",
		TitleName: "罗昂",
		TitleCode: "b_roanne",
	}
}
