package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法奥FaoBarony struct {
	feud.BaseBarony
}

var BFao法奥 feud.Barony = &法奥FaoBarony{}

func init() {
    f := BFao法奥.(*法奥FaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fao",
		TitleName: "法奥",
		TitleCode: "b_fao",
	}
}
