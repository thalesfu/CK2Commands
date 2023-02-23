package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班堡BamburghBarony struct {
	feud.BaseBarony
}

var BBamburgh班堡 feud.Barony = &班堡BamburghBarony{}

func init() {
	f := BBamburgh班堡.(*班堡BamburghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bamburgh",
		TitleName: "班堡",
		TitleCode: "b_bamburgh",
	}
}
