package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉科皮亚MalakopeiaBarony struct {
	feud.BaseBarony
}

var BMalakopeia马拉科皮亚 feud.Barony = &马拉科皮亚MalakopeiaBarony{}

func init() {
    f := BMalakopeia马拉科皮亚.(*马拉科皮亚MalakopeiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malakopeia",
		TitleName: "马拉科皮亚",
		TitleCode: "b_malakopeia",
	}
}
