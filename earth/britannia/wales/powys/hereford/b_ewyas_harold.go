package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤耶斯哈罗德Ewyas_haroldBarony struct {
	feud.BaseBarony
}

var BEwyas_harold尤耶斯哈罗德 feud.Barony = &尤耶斯哈罗德Ewyas_haroldBarony{}

func init() {
    f := BEwyas_harold尤耶斯哈罗德.(*尤耶斯哈罗德Ewyas_haroldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ewyas_harold",
		TitleName: "尤耶斯哈罗德",
		TitleCode: "b_ewyas_harold",
	}
}
