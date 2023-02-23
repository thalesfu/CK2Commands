package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞特SaterBarony struct {
	feud.BaseBarony
}

var BSater塞特 feud.Barony = &塞特SaterBarony{}

func init() {
	f := BSater塞特.(*塞特SaterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sater",
		TitleName: "塞特",
		TitleCode: "b_sater",
	}
}
