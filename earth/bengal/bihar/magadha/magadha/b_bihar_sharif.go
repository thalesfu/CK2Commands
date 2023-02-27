package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗诃罗谢里夫Bihar_sharifBarony struct {
	feud.BaseBarony
}

var BBihar_sharif毗诃罗谢里夫 feud.Barony = &毗诃罗谢里夫Bihar_sharifBarony{}

func init() {
    f := BBihar_sharif毗诃罗谢里夫.(*毗诃罗谢里夫Bihar_sharifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bihar_sharif",
		TitleName: "毗诃罗谢里夫",
		TitleCode: "b_bihar_sharif",
	}
}
