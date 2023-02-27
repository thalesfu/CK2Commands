package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼荼毗耶补罗MandavyapuraBarony struct {
	feud.BaseBarony
}

var BMandavyapura曼荼毗耶补罗 feud.Barony = &曼荼毗耶补罗MandavyapuraBarony{}

func init() {
    f := BMandavyapura曼荼毗耶补罗.(*曼荼毗耶补罗MandavyapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandavyapura",
		TitleName: "曼荼毗耶补罗",
		TitleCode: "b_mandavyapura",
	}
}
