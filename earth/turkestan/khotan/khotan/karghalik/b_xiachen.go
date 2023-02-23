package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏臣XiachenBarony struct {
	feud.BaseBarony
}

var BXiachen夏臣 feud.Barony = &夏臣XiachenBarony{}

func init() {
	f := BXiachen夏臣.(*夏臣XiachenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiachen",
		TitleName: "夏臣",
		TitleCode: "b_xiachen",
	}
}
