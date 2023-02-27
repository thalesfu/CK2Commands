package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷甘RegganeBarony struct {
	feud.BaseBarony
}

var BReggane雷甘 feud.Barony = &雷甘RegganeBarony{}

func init() {
    f := BReggane雷甘.(*雷甘RegganeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reggane",
		TitleName: "雷甘",
		TitleCode: "b_reggane",
	}
}
