package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴纳BanaBarony struct {
	feud.BaseBarony
}

var BBana巴纳 feud.Barony = &巴纳BanaBarony{}

func init() {
    f := BBana巴纳.(*巴纳BanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bana",
		TitleName: "巴纳",
		TitleCode: "b_bana",
	}
}
