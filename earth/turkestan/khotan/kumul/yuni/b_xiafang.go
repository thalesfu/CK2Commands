package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 狭方XiafangBarony struct {
	feud.BaseBarony
}

var BXiafang狭方 feud.Barony = &狭方XiafangBarony{}

func init() {
    f := BXiafang狭方.(*狭方XiafangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiafang",
		TitleName: "狭方",
		TitleCode: "b_xiafang",
	}
}
