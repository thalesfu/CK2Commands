package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴姆巴斯BambasiBarony struct {
	feud.BaseBarony
}

var BBambasi巴姆巴斯 feud.Barony = &巴姆巴斯BambasiBarony{}

func init() {
	f := BBambasi巴姆巴斯.(*巴姆巴斯BambasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bambasi",
		TitleName: "巴姆巴斯",
		TitleCode: "b_bambasi",
	}
}
