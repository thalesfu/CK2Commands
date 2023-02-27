package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特巴拉AtbaraBarony struct {
	feud.BaseBarony
}

var BAtbara阿特巴拉 feud.Barony = &阿特巴拉AtbaraBarony{}

func init() {
    f := BAtbara阿特巴拉.(*阿特巴拉AtbaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atbara",
		TitleName: "阿特巴拉",
		TitleCode: "b_atbara",
	}
}
