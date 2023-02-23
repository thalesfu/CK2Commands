package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶婆摩阇迦YavamajjhakaBarony struct {
	feud.BaseBarony
}

var BYavamajjhaka耶婆摩阇迦 feud.Barony = &耶婆摩阇迦YavamajjhakaBarony{}

func init() {
	f := BYavamajjhaka耶婆摩阇迦.(*耶婆摩阇迦YavamajjhakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yavamajjhaka",
		TitleName: "耶婆摩阇迦",
		TitleCode: "b_yavamajjhaka",
	}
}
