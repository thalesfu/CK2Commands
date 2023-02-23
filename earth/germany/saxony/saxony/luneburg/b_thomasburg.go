package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托马斯堡ThomasburgBarony struct {
	feud.BaseBarony
}

var BThomasburg托马斯堡 feud.Barony = &托马斯堡ThomasburgBarony{}

func init() {
	f := BThomasburg托马斯堡.(*托马斯堡ThomasburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thomasburg",
		TitleName: "托马斯堡",
		TitleCode: "b_thomasburg",
	}
}
