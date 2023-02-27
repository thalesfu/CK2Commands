package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布路沙布逻PurushapuraBarony struct {
	feud.BaseBarony
}

var BPurushapura布路沙布逻 feud.Barony = &布路沙布逻PurushapuraBarony{}

func init() {
    f := BPurushapura布路沙布逻.(*布路沙布逻PurushapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purushapura",
		TitleName: "布路沙布逻",
		TitleCode: "b_purushapura",
	}
}
