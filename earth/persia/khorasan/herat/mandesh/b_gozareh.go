package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古扎雷GozarehBarony struct {
	feud.BaseBarony
}

var BGozareh古扎雷 feud.Barony = &古扎雷GozarehBarony{}

func init() {
    f := BGozareh古扎雷.(*古扎雷GozarehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gozareh",
		TitleName: "古扎雷",
		TitleCode: "b_gozareh",
	}
}
