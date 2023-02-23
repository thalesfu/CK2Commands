package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班拉安BuraanBarony struct {
	feud.BaseBarony
}

var BBuraan班拉安 feud.Barony = &班拉安BuraanBarony{}

func init() {
	f := BBuraan班拉安.(*班拉安BuraanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buraan",
		TitleName: "班拉安",
		TitleCode: "b_buraan",
	}
}
