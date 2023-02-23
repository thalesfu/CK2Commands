package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎青ZhaqingBarony struct {
	feud.BaseBarony
}

var BZhaqing扎青 feud.Barony = &扎青ZhaqingBarony{}

func init() {
	f := BZhaqing扎青.(*扎青ZhaqingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhaqing",
		TitleName: "扎青",
		TitleCode: "b_zhaqing",
	}
}
