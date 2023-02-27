package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利兹LeedsBarony struct {
	feud.BaseBarony
}

var BLeeds利兹 feud.Barony = &利兹LeedsBarony{}

func init() {
    f := BLeeds利兹.(*利兹LeedsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leeds",
		TitleName: "利兹",
		TitleCode: "b_leeds",
	}
}
