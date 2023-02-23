package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温泉WenquanBarony struct {
	feud.BaseBarony
}

var BWenquan温泉 feud.Barony = &温泉WenquanBarony{}

func init() {
	f := BWenquan温泉.(*温泉WenquanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wenquan",
		TitleName: "温泉",
		TitleCode: "b_wenquan",
	}
}
