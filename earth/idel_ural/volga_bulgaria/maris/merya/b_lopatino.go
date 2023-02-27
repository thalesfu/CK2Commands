package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛帕季诺LopatinoBarony struct {
	feud.BaseBarony
}

var BLopatino洛帕季诺 feud.Barony = &洛帕季诺LopatinoBarony{}

func init() {
    f := BLopatino洛帕季诺.(*洛帕季诺LopatinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lopatino",
		TitleName: "洛帕季诺",
		TitleCode: "b_lopatino",
	}
}
