package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎波罗热ZaporzhyeBarony struct {
	feud.BaseBarony
}

var BZaporzhye扎波罗热 feud.Barony = &扎波罗热ZaporzhyeBarony{}

func init() {
    f := BZaporzhye扎波罗热.(*扎波罗热ZaporzhyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaporzhye",
		TitleName: "扎波罗热",
		TitleCode: "b_zaporzhye",
	}
}
