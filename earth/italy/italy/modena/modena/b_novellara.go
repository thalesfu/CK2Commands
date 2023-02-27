package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺韦拉拉NovellaraBarony struct {
	feud.BaseBarony
}

var BNovellara诺韦拉拉 feud.Barony = &诺韦拉拉NovellaraBarony{}

func init() {
    f := BNovellara诺韦拉拉.(*诺韦拉拉NovellaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novellara",
		TitleName: "诺韦拉拉",
		TitleCode: "b_novellara",
	}
}
