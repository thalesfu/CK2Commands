package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒克罗图瓦Le_crotoyBarony struct {
	feud.BaseBarony
}

var BLe_crotoy勒克罗图瓦 feud.Barony = &勒克罗图瓦Le_crotoyBarony{}

func init() {
    f := BLe_crotoy勒克罗图瓦.(*勒克罗图瓦Le_crotoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "le_crotoy",
		TitleName: "勒克罗图瓦",
		TitleCode: "b_le_crotoy",
	}
}
