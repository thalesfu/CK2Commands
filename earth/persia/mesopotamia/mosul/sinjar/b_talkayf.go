package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰勒凯夫TalkayfBarony struct {
	feud.BaseBarony
}

var BTalkayf泰勒凯夫 feud.Barony = &泰勒凯夫TalkayfBarony{}

func init() {
    f := BTalkayf泰勒凯夫.(*泰勒凯夫TalkayfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talkayf",
		TitleName: "泰勒凯夫",
		TitleCode: "b_talkayf",
	}
}
