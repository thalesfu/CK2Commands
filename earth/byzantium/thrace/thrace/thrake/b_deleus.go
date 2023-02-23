package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蔽日DeleusBarony struct {
	feud.BaseBarony
}

var BDeleus蔽日 feud.Barony = &蔽日DeleusBarony{}

func init() {
	f := BDeleus蔽日.(*蔽日DeleusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deleus",
		TitleName: "蔽日",
		TitleCode: "b_deleus",
	}
}
