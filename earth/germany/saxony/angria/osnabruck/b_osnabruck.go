package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯纳布吕克OsnabruckBarony struct {
	feud.BaseBarony
}

var BOsnabruck奥斯纳布吕克 feud.Barony = &奥斯纳布吕克OsnabruckBarony{}

func init() {
    f := BOsnabruck奥斯纳布吕克.(*奥斯纳布吕克OsnabruckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osnabruck",
		TitleName: "奥斯纳布吕克",
		TitleCode: "b_osnabruck",
	}
}
