package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯瓦特霍尔姆SvartholmBarony struct {
	feud.BaseBarony
}

var BSvartholm斯瓦特霍尔姆 feud.Barony = &斯瓦特霍尔姆SvartholmBarony{}

func init() {
	f := BSvartholm斯瓦特霍尔姆.(*斯瓦特霍尔姆SvartholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svartholm",
		TitleName: "斯瓦特霍尔姆",
		TitleCode: "b_svartholm",
	}
}
