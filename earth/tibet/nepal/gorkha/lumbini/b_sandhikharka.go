package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑提迦伽SandhikharkaBarony struct {
	feud.BaseBarony
}

var BSandhikharka桑提迦伽 feud.Barony = &桑提迦伽SandhikharkaBarony{}

func init() {
	f := BSandhikharka桑提迦伽.(*桑提迦伽SandhikharkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandhikharka",
		TitleName: "桑提迦伽",
		TitleCode: "b_sandhikharka",
	}
}
