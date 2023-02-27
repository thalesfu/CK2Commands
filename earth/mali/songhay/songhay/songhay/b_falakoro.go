package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉科罗FalakoroBarony struct {
	feud.BaseBarony
}

var BFalakoro法拉科罗 feud.Barony = &法拉科罗FalakoroBarony{}

func init() {
    f := BFalakoro法拉科罗.(*法拉科罗FalakoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falakoro",
		TitleName: "法拉科罗",
		TitleCode: "b_falakoro",
	}
}
