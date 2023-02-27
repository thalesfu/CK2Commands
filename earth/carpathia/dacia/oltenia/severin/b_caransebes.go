package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡兰塞贝什CaransebesBarony struct {
	feud.BaseBarony
}

var BCaransebes卡兰塞贝什 feud.Barony = &卡兰塞贝什CaransebesBarony{}

func init() {
    f := BCaransebes卡兰塞贝什.(*卡兰塞贝什CaransebesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caransebes",
		TitleName: "卡兰塞贝什",
		TitleCode: "b_caransebes",
	}
}
