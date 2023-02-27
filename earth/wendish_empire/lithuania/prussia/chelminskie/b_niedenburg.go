package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼登贝格NiedenburgBarony struct {
	feud.BaseBarony
}

var BNiedenburg尼登贝格 feud.Barony = &尼登贝格NiedenburgBarony{}

func init() {
    f := BNiedenburg尼登贝格.(*尼登贝格NiedenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niedenburg",
		TitleName: "尼登贝格",
		TitleCode: "b_niedenburg",
	}
}
