package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下斯图比察DonjastubicaBarony struct {
	feud.BaseBarony
}

var BDonjastubica下斯图比察 feud.Barony = &下斯图比察DonjastubicaBarony{}

func init() {
	f := BDonjastubica下斯图比察.(*下斯图比察DonjastubicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donjastubica",
		TitleName: "下斯图比察",
		TitleCode: "b_donjastubica",
	}
}
