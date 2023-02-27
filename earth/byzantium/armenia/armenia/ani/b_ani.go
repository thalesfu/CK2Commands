package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尼AniBarony struct {
	feud.BaseBarony
}

var BAni阿尼 feud.Barony = &阿尼AniBarony{}

func init() {
    f := BAni阿尼.(*阿尼AniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ani",
		TitleName: "阿尼",
		TitleCode: "b_ani",
	}
}
