package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯普隆格CampulungBarony struct {
	feud.BaseBarony
}

var BCampulung肯普隆格 feud.Barony = &肯普隆格CampulungBarony{}

func init() {
	f := BCampulung肯普隆格.(*肯普隆格CampulungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "campulung",
		TitleName: "肯普隆格",
		TitleCode: "b_campulung",
	}
}
