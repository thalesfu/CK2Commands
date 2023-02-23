package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米拉贝尔MirabelBarony struct {
	feud.BaseBarony
}

var BMirabel米拉贝尔 feud.Barony = &米拉贝尔MirabelBarony{}

func init() {
	f := BMirabel米拉贝尔.(*米拉贝尔MirabelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirabel",
		TitleName: "米拉贝尔",
		TitleCode: "b_mirabel",
	}
}
