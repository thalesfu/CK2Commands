package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑摩丽湿伐利SamaleswariBarony struct {
	feud.BaseBarony
}

var BSamaleswari娑摩丽湿伐利 feud.Barony = &娑摩丽湿伐利SamaleswariBarony{}

func init() {
    f := BSamaleswari娑摩丽湿伐利.(*娑摩丽湿伐利SamaleswariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samaleswari",
		TitleName: "娑摩丽湿伐利",
		TitleCode: "b_samaleswari",
	}
}
