package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热列兹尼亚基ZheleznyakiBarony struct {
	feud.BaseBarony
}

var BZheleznyaki热列兹尼亚基 feud.Barony = &热列兹尼亚基ZheleznyakiBarony{}

func init() {
	f := BZheleznyaki热列兹尼亚基.(*热列兹尼亚基ZheleznyakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zheleznyaki",
		TitleName: "热列兹尼亚基",
		TitleCode: "b_zheleznyaki",
	}
}
