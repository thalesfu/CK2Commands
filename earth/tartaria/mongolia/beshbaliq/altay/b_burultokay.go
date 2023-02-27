package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布伦托海BurultokayBarony struct {
	feud.BaseBarony
}

var BBurultokay布伦托海 feud.Barony = &布伦托海BurultokayBarony{}

func init() {
    f := BBurultokay布伦托海.(*布伦托海BurultokayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burultokay",
		TitleName: "布伦托海",
		TitleCode: "b_burultokay",
	}
}
