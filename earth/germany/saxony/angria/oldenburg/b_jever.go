package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶弗尔JeverBarony struct {
	feud.BaseBarony
}

var BJever耶弗尔 feud.Barony = &耶弗尔JeverBarony{}

func init() {
    f := BJever耶弗尔.(*耶弗尔JeverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jever",
		TitleName: "耶弗尔",
		TitleCode: "b_jever",
	}
}
