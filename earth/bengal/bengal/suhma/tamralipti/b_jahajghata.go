package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶诃耶迦陀JahajghataBarony struct {
	feud.BaseBarony
}

var BJahajghata耶诃耶迦陀 feud.Barony = &耶诃耶迦陀JahajghataBarony{}

func init() {
    f := BJahajghata耶诃耶迦陀.(*耶诃耶迦陀JahajghataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jahajghata",
		TitleName: "耶诃耶迦陀",
		TitleCode: "b_jahajghata",
	}
}
