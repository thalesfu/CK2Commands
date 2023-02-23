package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁卢斯BurlusBarony struct {
	feud.BaseBarony
}

var BBurlus布鲁卢斯 feud.Barony = &布鲁卢斯BurlusBarony{}

func init() {
	f := BBurlus布鲁卢斯.(*布鲁卢斯BurlusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burlus",
		TitleName: "布鲁卢斯",
		TitleCode: "b_burlus",
	}
}
