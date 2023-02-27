package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔布都奇DaarbuduqBarony struct {
	feud.BaseBarony
}

var BDaarbuduq达尔布都奇 feud.Barony = &达尔布都奇DaarbuduqBarony{}

func init() {
    f := BDaarbuduq达尔布都奇.(*达尔布都奇DaarbuduqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daarbuduq",
		TitleName: "达尔布都奇",
		TitleCode: "b_daarbuduq",
	}
}
