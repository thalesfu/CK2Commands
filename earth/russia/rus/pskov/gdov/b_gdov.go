package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格多夫GdovBarony struct {
	feud.BaseBarony
}

var BGdov格多夫 feud.Barony = &格多夫GdovBarony{}

func init() {
    f := BGdov格多夫.(*格多夫GdovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gdov",
		TitleName: "格多夫",
		TitleCode: "b_gdov",
	}
}
