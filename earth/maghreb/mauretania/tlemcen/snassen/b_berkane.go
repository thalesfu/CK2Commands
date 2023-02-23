package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜尔坎BerkaneBarony struct {
	feud.BaseBarony
}

var BBerkane拜尔坎 feud.Barony = &拜尔坎BerkaneBarony{}

func init() {
	f := BBerkane拜尔坎.(*拜尔坎BerkaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berkane",
		TitleName: "拜尔坎",
		TitleCode: "b_berkane",
	}
}
