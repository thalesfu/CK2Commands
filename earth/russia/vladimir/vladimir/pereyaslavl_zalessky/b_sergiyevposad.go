package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢尔吉耶夫波萨德SergiyevposadBarony struct {
	feud.BaseBarony
}

var BSergiyevposad谢尔吉耶夫波萨德 feud.Barony = &谢尔吉耶夫波萨德SergiyevposadBarony{}

func init() {
    f := BSergiyevposad谢尔吉耶夫波萨德.(*谢尔吉耶夫波萨德SergiyevposadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sergiyevposad",
		TitleName: "谢尔吉耶夫波萨德",
		TitleCode: "b_sergiyevposad",
	}
}
