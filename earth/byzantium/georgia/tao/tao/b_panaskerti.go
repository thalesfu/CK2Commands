package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕纳斯克尔蒂PanaskertiBarony struct {
	feud.BaseBarony
}

var BPanaskerti帕纳斯克尔蒂 feud.Barony = &帕纳斯克尔蒂PanaskertiBarony{}

func init() {
    f := BPanaskerti帕纳斯克尔蒂.(*帕纳斯克尔蒂PanaskertiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panaskerti",
		TitleName: "帕纳斯克尔蒂",
		TitleCode: "b_panaskerti",
	}
}
