package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲塞FuisseBarony struct {
	feud.BaseBarony
}

var BFuisse菲塞 feud.Barony = &菲塞FuisseBarony{}

func init() {
    f := BFuisse菲塞.(*菲塞FuisseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuisse",
		TitleName: "菲塞",
		TitleCode: "b_fuisse",
	}
}
