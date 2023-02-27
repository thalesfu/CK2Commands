package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌克梅尔盖UkmergeBarony struct {
	feud.BaseBarony
}

var BUkmerge乌克梅尔盖 feud.Barony = &乌克梅尔盖UkmergeBarony{}

func init() {
    f := BUkmerge乌克梅尔盖.(*乌克梅尔盖UkmergeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ukmerge",
		TitleName: "乌克梅尔盖",
		TitleCode: "b_ukmerge",
	}
}
