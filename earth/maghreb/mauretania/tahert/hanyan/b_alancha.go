package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃安查AlanchaBarony struct {
	feud.BaseBarony
}

var BAlancha埃安查 feud.Barony = &埃安查AlanchaBarony{}

func init() {
	f := BAlancha埃安查.(*埃安查AlanchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alancha",
		TitleName: "埃安查",
		TitleCode: "b_alancha",
	}
}
