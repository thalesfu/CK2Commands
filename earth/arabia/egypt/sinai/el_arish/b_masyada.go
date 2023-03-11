package el_arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马察达MasyadaBarony struct {
	feud.BaseBarony
}

var BMasyada马察达 feud.Barony = &马察达MasyadaBarony{}

func init() {
    f := BMasyada马察达.(*马察达MasyadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masyada",
		TitleName: "马察达",
		TitleCode: "b_masyada",
	}
}
