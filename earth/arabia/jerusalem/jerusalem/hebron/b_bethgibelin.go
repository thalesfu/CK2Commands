package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝特吉伯兰BethgibelinBarony struct {
	feud.BaseBarony
}

var BBethgibelin贝特吉伯兰 feud.Barony = &贝特吉伯兰BethgibelinBarony{}

func init() {
    f := BBethgibelin贝特吉伯兰.(*贝特吉伯兰BethgibelinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bethgibelin",
		TitleName: "贝特吉伯兰",
		TitleCode: "b_bethgibelin",
	}
}
