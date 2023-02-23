package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢列达SeredaBarony struct {
	feud.BaseBarony
}

var BSereda谢列达 feud.Barony = &谢列达SeredaBarony{}

func init() {
	f := BSereda谢列达.(*谢列达SeredaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sereda",
		TitleName: "谢列达",
		TitleCode: "b_sereda",
	}
}
