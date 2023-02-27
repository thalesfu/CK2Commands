package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉塔亚克KhratayakBarony struct {
	feud.BaseBarony
}

var BKhratayak赫拉塔亚克 feud.Barony = &赫拉塔亚克KhratayakBarony{}

func init() {
    f := BKhratayak赫拉塔亚克.(*赫拉塔亚克KhratayakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khratayak",
		TitleName: "赫拉塔亚克",
		TitleCode: "b_khratayak",
	}
}
