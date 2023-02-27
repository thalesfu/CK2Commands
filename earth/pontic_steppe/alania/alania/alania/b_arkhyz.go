package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔黑兹ArkhyzBarony struct {
	feud.BaseBarony
}

var BArkhyz阿尔黑兹 feud.Barony = &阿尔黑兹ArkhyzBarony{}

func init() {
    f := BArkhyz阿尔黑兹.(*阿尔黑兹ArkhyzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkhyz",
		TitleName: "阿尔黑兹",
		TitleCode: "b_arkhyz",
	}
}
