package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圭尔福堡CastelguelfoBarony struct {
	feud.BaseBarony
}

var BCastelguelfo圭尔福堡 feud.Barony = &圭尔福堡CastelguelfoBarony{}

func init() {
    f := BCastelguelfo圭尔福堡.(*圭尔福堡CastelguelfoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelguelfo",
		TitleName: "圭尔福堡",
		TitleCode: "b_castelguelfo",
	}
}
