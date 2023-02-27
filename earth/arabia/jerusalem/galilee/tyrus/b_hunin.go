package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡灵HuninBarony struct {
	feud.BaseBarony
}

var BHunin胡灵 feud.Barony = &胡灵HuninBarony{}

func init() {
    f := BHunin胡灵.(*胡灵HuninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hunin",
		TitleName: "胡灵",
		TitleCode: "b_hunin",
	}
}
