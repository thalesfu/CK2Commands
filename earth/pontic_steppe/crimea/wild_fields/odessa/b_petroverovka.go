package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得罗韦罗夫卡PetroverovkaBarony struct {
	feud.BaseBarony
}

var BPetroverovka彼得罗韦罗夫卡 feud.Barony = &彼得罗韦罗夫卡PetroverovkaBarony{}

func init() {
    f := BPetroverovka彼得罗韦罗夫卡.(*彼得罗韦罗夫卡PetroverovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petroverovka",
		TitleName: "彼得罗韦罗夫卡",
		TitleCode: "b_petroverovka",
	}
}
