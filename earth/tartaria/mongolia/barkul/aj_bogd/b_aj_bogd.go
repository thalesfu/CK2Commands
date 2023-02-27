package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿主孛黑答Aj_bogdBarony struct {
	feud.BaseBarony
}

var BAj_bogd阿主孛黑答 feud.Barony = &阿主孛黑答Aj_bogdBarony{}

func init() {
    f := BAj_bogd阿主孛黑答.(*阿主孛黑答Aj_bogdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aj_bogd",
		TitleName: "阿主孛黑答",
		TitleCode: "b_aj_bogd",
	}
}
