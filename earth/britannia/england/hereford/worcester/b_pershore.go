package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀肖尔PershoreBarony struct {
	feud.BaseBarony
}

var BPershore珀肖尔 feud.Barony = &珀肖尔PershoreBarony{}

func init() {
	f := BPershore珀肖尔.(*珀肖尔PershoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pershore",
		TitleName: "珀肖尔",
		TitleCode: "b_pershore",
	}
}
