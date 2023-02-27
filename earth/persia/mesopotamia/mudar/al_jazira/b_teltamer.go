package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特达梅尔TeltamerBarony struct {
	feud.BaseBarony
}

var BTeltamer特达梅尔 feud.Barony = &特达梅尔TeltamerBarony{}

func init() {
    f := BTeltamer特达梅尔.(*特达梅尔TeltamerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teltamer",
		TitleName: "特达梅尔",
		TitleCode: "b_teltamer",
	}
}
