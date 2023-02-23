package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉奎拉AquilaBarony struct {
	feud.BaseBarony
}

var BAquila拉奎拉 feud.Barony = &拉奎拉AquilaBarony{}

func init() {
	f := BAquila拉奎拉.(*拉奎拉AquilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aquila",
		TitleName: "拉奎拉",
		TitleCode: "b_aquila",
	}
}
