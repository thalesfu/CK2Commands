package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕雷德斯ParedesBarony struct {
	feud.BaseBarony
}

var BParedes帕雷德斯 feud.Barony = &帕雷德斯ParedesBarony{}

func init() {
    f := BParedes帕雷德斯.(*帕雷德斯ParedesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paredes",
		TitleName: "帕雷德斯",
		TitleCode: "b_paredes",
	}
}
