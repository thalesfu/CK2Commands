package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布罗姆斯格罗夫BromsgroveBarony struct {
	feud.BaseBarony
}

var BBromsgrove布罗姆斯格罗夫 feud.Barony = &布罗姆斯格罗夫BromsgroveBarony{}

func init() {
    f := BBromsgrove布罗姆斯格罗夫.(*布罗姆斯格罗夫BromsgroveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bromsgrove",
		TitleName: "布罗姆斯格罗夫",
		TitleCode: "b_bromsgrove",
	}
}
