package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦日戈尔特VazhgortBarony struct {
	feud.BaseBarony
}

var BVazhgort瓦日戈尔特 feud.Barony = &瓦日戈尔特VazhgortBarony{}

func init() {
    f := BVazhgort瓦日戈尔特.(*瓦日戈尔特VazhgortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vazhgort",
		TitleName: "瓦日戈尔特",
		TitleCode: "b_vazhgort",
	}
}
