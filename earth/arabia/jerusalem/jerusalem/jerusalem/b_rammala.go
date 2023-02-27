package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉马拉RammalaBarony struct {
	feud.BaseBarony
}

var BRammala拉马拉 feud.Barony = &拉马拉RammalaBarony{}

func init() {
    f := BRammala拉马拉.(*拉马拉RammalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rammala",
		TitleName: "拉马拉",
		TitleCode: "b_rammala",
	}
}
