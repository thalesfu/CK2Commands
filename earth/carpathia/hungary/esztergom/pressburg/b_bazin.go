package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍津BazinBarony struct {
	feud.BaseBarony
}

var BBazin鲍津 feud.Barony = &鲍津BazinBarony{}

func init() {
    f := BBazin鲍津.(*鲍津BazinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bazin",
		TitleName: "鲍津",
		TitleCode: "b_bazin",
	}
}
