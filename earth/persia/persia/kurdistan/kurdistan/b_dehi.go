package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代希DehiBarony struct {
	feud.BaseBarony
}

var BDehi代希 feud.Barony = &代希DehiBarony{}

func init() {
	f := BDehi代希.(*代希DehiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dehi",
		TitleName: "代希",
		TitleCode: "b_dehi",
	}
}
