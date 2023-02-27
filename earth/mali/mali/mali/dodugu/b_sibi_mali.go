package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西比Sibi_maliBarony struct {
	feud.BaseBarony
}

var BSibi_mali西比 feud.Barony = &西比Sibi_maliBarony{}

func init() {
    f := BSibi_mali西比.(*西比Sibi_maliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sibi_mali",
		TitleName: "西比",
		TitleCode: "b_sibi_mali",
	}
}
