package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜申泰BachantaBarony struct {
	feud.BaseBarony
}

var BBachanta拜申泰 feud.Barony = &拜申泰BachantaBarony{}

func init() {
    f := BBachanta拜申泰.(*拜申泰BachantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bachanta",
		TitleName: "拜申泰",
		TitleCode: "b_bachanta",
	}
}
