package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙谢拉NowsheraBarony struct {
	feud.BaseBarony
}

var BNowshera瑙谢拉 feud.Barony = &瑙谢拉NowsheraBarony{}

func init() {
    f := BNowshera瑙谢拉.(*瑙谢拉NowsheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nowshera",
		TitleName: "瑙谢拉",
		TitleCode: "b_nowshera",
	}
}
