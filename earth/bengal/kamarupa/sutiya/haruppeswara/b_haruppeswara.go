package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃卢毗湿伐罗HaruppeswaraBarony struct {
	feud.BaseBarony
}

var BHaruppeswara诃卢毗湿伐罗 feud.Barony = &诃卢毗湿伐罗HaruppeswaraBarony{}

func init() {
    f := BHaruppeswara诃卢毗湿伐罗.(*诃卢毗湿伐罗HaruppeswaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haruppeswara",
		TitleName: "诃卢毗湿伐罗",
		TitleCode: "b_haruppeswara",
	}
}
