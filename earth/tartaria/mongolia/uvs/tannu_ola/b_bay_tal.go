package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜塔尔Bay_talBarony struct {
	feud.BaseBarony
}

var BBay_tal拜塔尔 feud.Barony = &拜塔尔Bay_talBarony{}

func init() {
    f := BBay_tal拜塔尔.(*拜塔尔Bay_talBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bay_tal",
		TitleName: "拜塔尔",
		TitleCode: "b_bay_tal",
	}
}
