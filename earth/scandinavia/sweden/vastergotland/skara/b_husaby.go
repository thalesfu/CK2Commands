package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许萨比HusabyBarony struct {
	feud.BaseBarony
}

var BHusaby许萨比 feud.Barony = &许萨比HusabyBarony{}

func init() {
	f := BHusaby许萨比.(*许萨比HusabyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "husaby",
		TitleName: "许萨比",
		TitleCode: "b_husaby",
	}
}
