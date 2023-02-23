package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫雷利亚MorellaBarony struct {
	feud.BaseBarony
}

var BMorella莫雷利亚 feud.Barony = &莫雷利亚MorellaBarony{}

func init() {
	f := BMorella莫雷利亚.(*莫雷利亚MorellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morella",
		TitleName: "莫雷利亚",
		TitleCode: "b_morella",
	}
}
