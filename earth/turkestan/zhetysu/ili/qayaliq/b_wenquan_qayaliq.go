package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温泉Wenquan_qayaliqBarony struct {
	feud.BaseBarony
}

var BWenquan_qayaliq温泉 feud.Barony = &温泉Wenquan_qayaliqBarony{}

func init() {
    f := BWenquan_qayaliq温泉.(*温泉Wenquan_qayaliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wenquan_qayaliq",
		TitleName: "温泉",
		TitleCode: "b_wenquan_qayaliq",
	}
}
