package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特罗CastroBarony struct {
	feud.BaseBarony
}

var BCastro卡斯特罗 feud.Barony = &卡斯特罗CastroBarony{}

func init() {
	f := BCastro卡斯特罗.(*卡斯特罗CastroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castro",
		TitleName: "卡斯特罗",
		TitleCode: "b_castro",
	}
}
