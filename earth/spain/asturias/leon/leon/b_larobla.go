package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗夫拉LaroblaBarony struct {
	feud.BaseBarony
}

var BLarobla拉罗夫拉 feud.Barony = &拉罗夫拉LaroblaBarony{}

func init() {
    f := BLarobla拉罗夫拉.(*拉罗夫拉LaroblaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larobla",
		TitleName: "拉罗夫拉",
		TitleCode: "b_larobla",
	}
}
