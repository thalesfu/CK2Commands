package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下卡姆斯克NizhnekamskBarony struct {
	feud.BaseBarony
}

var BNizhnekamsk下卡姆斯克 feud.Barony = &下卡姆斯克NizhnekamskBarony{}

func init() {
    f := BNizhnekamsk下卡姆斯克.(*下卡姆斯克NizhnekamskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhnekamsk",
		TitleName: "下卡姆斯克",
		TitleCode: "b_nizhnekamsk",
	}
}
