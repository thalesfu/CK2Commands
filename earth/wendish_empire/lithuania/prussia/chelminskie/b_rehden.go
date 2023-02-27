package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷登RehdenBarony struct {
	feud.BaseBarony
}

var BRehden雷登 feud.Barony = &雷登RehdenBarony{}

func init() {
    f := BRehden雷登.(*雷登RehdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rehden",
		TitleName: "雷登",
		TitleCode: "b_rehden",
	}
}
