package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尸毗SibiBarony struct {
	feud.BaseBarony
}

var BSibi尸毗 feud.Barony = &尸毗SibiBarony{}

func init() {
    f := BSibi尸毗.(*尸毗SibiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sibi",
		TitleName: "尸毗",
		TitleCode: "b_sibi",
	}
}
