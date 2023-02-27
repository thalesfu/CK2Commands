package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加格GagBarony struct {
	feud.BaseBarony
}

var BGag加格 feud.Barony = &加格GagBarony{}

func init() {
    f := BGag加格.(*加格GagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gag",
		TitleName: "加格",
		TitleCode: "b_gag",
	}
}
