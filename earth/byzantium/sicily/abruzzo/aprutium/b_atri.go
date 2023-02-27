package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特里AtriBarony struct {
	feud.BaseBarony
}

var BAtri阿特里 feud.Barony = &阿特里AtriBarony{}

func init() {
    f := BAtri阿特里.(*阿特里AtriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atri",
		TitleName: "阿特里",
		TitleCode: "b_atri",
	}
}
