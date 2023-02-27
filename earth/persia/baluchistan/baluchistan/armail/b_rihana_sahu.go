package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蕾哈纳沙笏Rihana_sahuBarony struct {
	feud.BaseBarony
}

var BRihana_sahu蕾哈纳沙笏 feud.Barony = &蕾哈纳沙笏Rihana_sahuBarony{}

func init() {
    f := BRihana_sahu蕾哈纳沙笏.(*蕾哈纳沙笏Rihana_sahuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rihana_sahu",
		TitleName: "蕾哈纳沙笏",
		TitleCode: "b_rihana_sahu",
	}
}
