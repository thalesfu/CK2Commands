package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒米哈ShumikhaBarony struct {
	feud.BaseBarony
}

var BShumikha舒米哈 feud.Barony = &舒米哈ShumikhaBarony{}

func init() {
    f := BShumikha舒米哈.(*舒米哈ShumikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shumikha",
		TitleName: "舒米哈",
		TitleCode: "b_shumikha",
	}
}
