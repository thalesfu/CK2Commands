package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 津吉巴尔ZinjibarBarony struct {
	feud.BaseBarony
}

var BZinjibar津吉巴尔 feud.Barony = &津吉巴尔ZinjibarBarony{}

func init() {
    f := BZinjibar津吉巴尔.(*津吉巴尔ZinjibarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zinjibar",
		TitleName: "津吉巴尔",
		TitleCode: "b_zinjibar",
	}
}
