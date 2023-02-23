package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿殊檀AjodhanBarony struct {
	feud.BaseBarony
}

var BAjodhan阿殊檀 feud.Barony = &阿殊檀AjodhanBarony{}

func init() {
	f := BAjodhan阿殊檀.(*阿殊檀AjodhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajodhan",
		TitleName: "阿殊檀",
		TitleCode: "b_ajodhan",
	}
}
