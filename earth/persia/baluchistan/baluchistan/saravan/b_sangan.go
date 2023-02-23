package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑甘SanganBarony struct {
	feud.BaseBarony
}

var BSangan桑甘 feud.Barony = &桑甘SanganBarony{}

func init() {
	f := BSangan桑甘.(*桑甘SanganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangan",
		TitleName: "桑甘",
		TitleCode: "b_sangan",
	}
}
