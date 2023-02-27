package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉尤埃拉TalayuelaBarony struct {
	feud.BaseBarony
}

var BTalayuela塔拉尤埃拉 feud.Barony = &塔拉尤埃拉TalayuelaBarony{}

func init() {
    f := BTalayuela塔拉尤埃拉.(*塔拉尤埃拉TalayuelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talayuela",
		TitleName: "塔拉尤埃拉",
		TitleCode: "b_talayuela",
	}
}
