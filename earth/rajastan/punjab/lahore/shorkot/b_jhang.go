package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 章JhangBarony struct {
	feud.BaseBarony
}

var BJhang章 feud.Barony = &章JhangBarony{}

func init() {
	f := BJhang章.(*章JhangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhang",
		TitleName: "章",
		TitleCode: "b_jhang",
	}
}
