package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿谢_赛AshchesayBarony struct {
	feud.BaseBarony
}

var BAshchesay阿谢_赛 feud.Barony = &阿谢_赛AshchesayBarony{}

func init() {
    f := BAshchesay阿谢_赛.(*阿谢_赛AshchesayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashchesay",
		TitleName: "阿谢_赛",
		TitleCode: "b_ashchesay",
	}
}
