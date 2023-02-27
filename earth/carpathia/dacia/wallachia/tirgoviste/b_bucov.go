package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布科夫BucovBarony struct {
	feud.BaseBarony
}

var BBucov布科夫 feud.Barony = &布科夫BucovBarony{}

func init() {
    f := BBucov布科夫.(*布科夫BucovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bucov",
		TitleName: "布科夫",
		TitleCode: "b_bucov",
	}
}
