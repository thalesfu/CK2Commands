package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡文CavanBarony struct {
	feud.BaseBarony
}

var BCavan卡文 feud.Barony = &卡文CavanBarony{}

func init() {
	f := BCavan卡文.(*卡文CavanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cavan",
		TitleName: "卡文",
		TitleCode: "b_cavan",
	}
}
