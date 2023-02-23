package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯里斯特威斯AberystwythBarony struct {
	feud.BaseBarony
}

var BAberystwyth阿伯里斯特威斯 feud.Barony = &阿伯里斯特威斯AberystwythBarony{}

func init() {
	f := BAberystwyth阿伯里斯特威斯.(*阿伯里斯特威斯AberystwythBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aberystwyth",
		TitleName: "阿伯里斯特威斯",
		TitleCode: "b_aberystwyth",
	}
}
