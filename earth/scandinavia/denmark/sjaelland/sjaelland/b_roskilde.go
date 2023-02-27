package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯基勒RoskildeBarony struct {
	feud.BaseBarony
}

var BRoskilde罗斯基勒 feud.Barony = &罗斯基勒RoskildeBarony{}

func init() {
    f := BRoskilde罗斯基勒.(*罗斯基勒RoskildeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roskilde",
		TitleName: "罗斯基勒",
		TitleCode: "b_roskilde",
	}
}
