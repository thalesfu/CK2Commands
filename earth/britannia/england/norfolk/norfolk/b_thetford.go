package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞特福特ThetfordBarony struct {
	feud.BaseBarony
}

var BThetford塞特福特 feud.Barony = &塞特福特ThetfordBarony{}

func init() {
	f := BThetford塞特福特.(*塞特福特ThetfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thetford",
		TitleName: "塞特福特",
		TitleCode: "b_thetford",
	}
}
