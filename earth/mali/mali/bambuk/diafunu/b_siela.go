package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西埃拉SielaBarony struct {
	feud.BaseBarony
}

var BSiela西埃拉 feud.Barony = &西埃拉SielaBarony{}

func init() {
	f := BSiela西埃拉.(*西埃拉SielaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siela",
		TitleName: "西埃拉",
		TitleCode: "b_siela",
	}
}
