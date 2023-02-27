package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马安MaanBarony struct {
	feud.BaseBarony
}

var BMaan马安 feud.Barony = &马安MaanBarony{}

func init() {
    f := BMaan马安.(*马安MaanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maan",
		TitleName: "马安",
		TitleCode: "b_maan",
	}
}
