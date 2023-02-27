package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普雷斯堡PressburgBarony struct {
	feud.BaseBarony
}

var BPressburg普雷斯堡 feud.Barony = &普雷斯堡PressburgBarony{}

func init() {
    f := BPressburg普雷斯堡.(*普雷斯堡PressburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pressburg",
		TitleName: "普雷斯堡",
		TitleCode: "b_pressburg",
	}
}
