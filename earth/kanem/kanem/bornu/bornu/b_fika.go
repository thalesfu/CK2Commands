package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲卡FikaBarony struct {
	feud.BaseBarony
}

var BFika菲卡 feud.Barony = &菲卡FikaBarony{}

func init() {
    f := BFika菲卡.(*菲卡FikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fika",
		TitleName: "菲卡",
		TitleCode: "b_fika",
	}
}
