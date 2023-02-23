package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 堤亚纳TyanaBarony struct {
	feud.BaseBarony
}

var BTyana堤亚纳 feud.Barony = &堤亚纳TyanaBarony{}

func init() {
	f := BTyana堤亚纳.(*堤亚纳TyanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyana",
		TitleName: "堤亚纳",
		TitleCode: "b_tyana",
	}
}
