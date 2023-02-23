package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡代诺SidernoBarony struct {
	feud.BaseBarony
}

var BSiderno锡代诺 feud.Barony = &锡代诺SidernoBarony{}

func init() {
	f := BSiderno锡代诺.(*锡代诺SidernoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siderno",
		TitleName: "锡代诺",
		TitleCode: "b_siderno",
	}
}
