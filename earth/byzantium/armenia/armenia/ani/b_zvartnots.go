package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹瓦尔特诺茨ZvartnotsBarony struct {
	feud.BaseBarony
}

var BZvartnots兹瓦尔特诺茨 feud.Barony = &兹瓦尔特诺茨ZvartnotsBarony{}

func init() {
    f := BZvartnots兹瓦尔特诺茨.(*兹瓦尔特诺茨ZvartnotsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zvartnots",
		TitleName: "兹瓦尔特诺茨",
		TitleCode: "b_zvartnots",
	}
}
