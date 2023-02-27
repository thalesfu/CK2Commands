package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布基斯齐赫BukistsikheBarony struct {
	feud.BaseBarony
}

var BBukistsikhe布基斯齐赫 feud.Barony = &布基斯齐赫BukistsikheBarony{}

func init() {
    f := BBukistsikhe布基斯齐赫.(*布基斯齐赫BukistsikheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bukistsikhe",
		TitleName: "布基斯齐赫",
		TitleCode: "b_bukistsikhe",
	}
}
