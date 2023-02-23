package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旁揭罗BangraBarony struct {
	feud.BaseBarony
}

var BBangra旁揭罗 feud.Barony = &旁揭罗BangraBarony{}

func init() {
	f := BBangra旁揭罗.(*旁揭罗BangraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangra",
		TitleName: "旁揭罗",
		TitleCode: "b_bangra",
	}
}
