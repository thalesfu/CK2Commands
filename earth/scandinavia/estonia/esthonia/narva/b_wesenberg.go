package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦森贝格WesenbergBarony struct {
	feud.BaseBarony
}

var BWesenberg韦森贝格 feud.Barony = &韦森贝格WesenbergBarony{}

func init() {
	f := BWesenberg韦森贝格.(*韦森贝格WesenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wesenberg",
		TitleName: "韦森贝格",
		TitleCode: "b_wesenberg",
	}
}
