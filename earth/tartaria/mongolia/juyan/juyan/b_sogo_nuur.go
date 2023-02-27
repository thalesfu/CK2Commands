package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索果淖尔Sogo_nuurBarony struct {
	feud.BaseBarony
}

var BSogo_nuur索果淖尔 feud.Barony = &索果淖尔Sogo_nuurBarony{}

func init() {
    f := BSogo_nuur索果淖尔.(*索果淖尔Sogo_nuurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sogo_nuur",
		TitleName: "索果淖尔",
		TitleCode: "b_sogo_nuur",
	}
}
