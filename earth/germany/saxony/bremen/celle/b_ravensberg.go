package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉芬斯堡RavensbergBarony struct {
	feud.BaseBarony
}

var BRavensberg拉芬斯堡 feud.Barony = &拉芬斯堡RavensbergBarony{}

func init() {
    f := BRavensberg拉芬斯堡.(*拉芬斯堡RavensbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravensberg",
		TitleName: "拉芬斯堡",
		TitleCode: "b_ravensberg",
	}
}
