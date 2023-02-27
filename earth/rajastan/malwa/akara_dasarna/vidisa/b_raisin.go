package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗僧RaisinBarony struct {
	feud.BaseBarony
}

var BRaisin罗僧 feud.Barony = &罗僧RaisinBarony{}

func init() {
    f := BRaisin罗僧.(*罗僧RaisinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raisin",
		TitleName: "罗僧",
		TitleCode: "b_raisin",
	}
}
