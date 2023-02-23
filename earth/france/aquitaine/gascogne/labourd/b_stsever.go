package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣瑟韦StseverBarony struct {
	feud.BaseBarony
}

var BStsever圣瑟韦 feud.Barony = &圣瑟韦StseverBarony{}

func init() {
	f := BStsever圣瑟韦.(*圣瑟韦StseverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stsever",
		TitleName: "圣瑟韦",
		TitleCode: "b_stsever",
	}
}
