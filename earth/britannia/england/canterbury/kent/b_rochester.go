package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗切斯特RochesterBarony struct {
	feud.BaseBarony
}

var BRochester罗切斯特 feud.Barony = &罗切斯特RochesterBarony{}

func init() {
    f := BRochester罗切斯特.(*罗切斯特RochesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rochester",
		TitleName: "罗切斯特",
		TitleCode: "b_rochester",
	}
}
