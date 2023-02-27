package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴内伊BanyyBarony struct {
	feud.BaseBarony
}

var BBanyy巴内伊 feud.Barony = &巴内伊BanyyBarony{}

func init() {
    f := BBanyy巴内伊.(*巴内伊BanyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banyy",
		TitleName: "巴内伊",
		TitleCode: "b_banyy",
	}
}
