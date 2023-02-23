package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫云MoyunBarony struct {
	feud.BaseBarony
}

var BMoyun莫云 feud.Barony = &莫云MoyunBarony{}

func init() {
	f := BMoyun莫云.(*莫云MoyunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moyun",
		TitleName: "莫云",
		TitleCode: "b_moyun",
	}
}
