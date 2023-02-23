package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫利赞RelizaneBarony struct {
	feud.BaseBarony
}

var BRelizane赫利赞 feud.Barony = &赫利赞RelizaneBarony{}

func init() {
	f := BRelizane赫利赞.(*赫利赞RelizaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "relizane",
		TitleName: "赫利赞",
		TitleCode: "b_relizane",
	}
}
