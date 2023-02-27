package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 荼稽湿伐罗阇底耶神庙Dhakeshwari_jatiya_mandirBarony struct {
	feud.BaseBarony
}

var BDhakeshwari_jatiya_mandir荼稽湿伐罗阇底耶神庙 feud.Barony = &荼稽湿伐罗阇底耶神庙Dhakeshwari_jatiya_mandirBarony{}

func init() {
    f := BDhakeshwari_jatiya_mandir荼稽湿伐罗阇底耶神庙.(*荼稽湿伐罗阇底耶神庙Dhakeshwari_jatiya_mandirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhakeshwari_jatiya_mandir",
		TitleName: "荼稽湿伐罗阇底耶神庙",
		TitleCode: "b_dhakeshwari_jatiya_mandir",
	}
}
