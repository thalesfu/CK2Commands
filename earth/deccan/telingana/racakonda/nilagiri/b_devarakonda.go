package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆罗军荼DevarakondaBarony struct {
	feud.BaseBarony
}

var BDevarakonda提婆罗军荼 feud.Barony = &提婆罗军荼DevarakondaBarony{}

func init() {
	f := BDevarakonda提婆罗军荼.(*提婆罗军荼DevarakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devarakonda",
		TitleName: "提婆罗军荼",
		TitleCode: "b_devarakonda",
	}
}
