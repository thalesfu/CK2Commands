package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔耶堡TaillebourgBarony struct {
	feud.BaseBarony
}

var BTaillebourg塔耶堡 feud.Barony = &塔耶堡TaillebourgBarony{}

func init() {
    f := BTaillebourg塔耶堡.(*塔耶堡TaillebourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taillebourg",
		TitleName: "塔耶堡",
		TitleCode: "b_taillebourg",
	}
}
