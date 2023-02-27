package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗奈RosnayBarony struct {
	feud.BaseBarony
}

var BRosnay罗奈 feud.Barony = &罗奈RosnayBarony{}

func init() {
    f := BRosnay罗奈.(*罗奈RosnayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rosnay",
		TitleName: "罗奈",
		TitleCode: "b_rosnay",
	}
}
