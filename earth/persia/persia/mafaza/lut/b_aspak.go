package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯帕克AspakBarony struct {
	feud.BaseBarony
}

var BAspak阿斯帕克 feud.Barony = &阿斯帕克AspakBarony{}

func init() {
    f := BAspak阿斯帕克.(*阿斯帕克AspakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aspak",
		TitleName: "阿斯帕克",
		TitleCode: "b_aspak",
	}
}
