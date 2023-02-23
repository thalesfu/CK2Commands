package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑布鲁SambourouBarony struct {
	feud.BaseBarony
}

var BSambourou桑布鲁 feud.Barony = &桑布鲁SambourouBarony{}

func init() {
	f := BSambourou桑布鲁.(*桑布鲁SambourouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sambourou",
		TitleName: "桑布鲁",
		TitleCode: "b_sambourou",
	}
}
