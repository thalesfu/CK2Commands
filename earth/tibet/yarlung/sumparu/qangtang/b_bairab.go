package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜惹布BairabBarony struct {
	feud.BaseBarony
}

var BBairab拜惹布 feud.Barony = &拜惹布BairabBarony{}

func init() {
	f := BBairab拜惹布.(*拜惹布BairabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bairab",
		TitleName: "拜惹布",
		TitleCode: "b_bairab",
	}
}
