package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_杰古塔UstdzeghutaBarony struct {
	feud.BaseBarony
}

var BUstdzeghuta乌斯季_杰古塔 feud.Barony = &乌斯季_杰古塔UstdzeghutaBarony{}

func init() {
    f := BUstdzeghuta乌斯季_杰古塔.(*乌斯季_杰古塔UstdzeghutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustdzeghuta",
		TitleName: "乌斯季_杰古塔",
		TitleCode: "b_ustdzeghuta",
	}
}
