package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢菲尔德SheffieldBarony struct {
	feud.BaseBarony
}

var BSheffield谢菲尔德 feud.Barony = &谢菲尔德SheffieldBarony{}

func init() {
    f := BSheffield谢菲尔德.(*谢菲尔德SheffieldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sheffield",
		TitleName: "谢菲尔德",
		TitleCode: "b_sheffield",
	}
}
