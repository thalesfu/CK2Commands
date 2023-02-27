package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌雷尔UrylBarony struct {
	feud.BaseBarony
}

var BUryl乌雷尔 feud.Barony = &乌雷尔UrylBarony{}

func init() {
    f := BUryl乌雷尔.(*乌雷尔UrylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uryl",
		TitleName: "乌雷尔",
		TitleCode: "b_uryl",
	}
}
