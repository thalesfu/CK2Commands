package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯蒂亚BastiaBarony struct {
	feud.BaseBarony
}

var BBastia巴斯蒂亚 feud.Barony = &巴斯蒂亚BastiaBarony{}

func init() {
	f := BBastia巴斯蒂亚.(*巴斯蒂亚BastiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bastia",
		TitleName: "巴斯蒂亚",
		TitleCode: "b_bastia",
	}
}
