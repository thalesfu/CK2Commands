package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比哈奇BihacBarony struct {
	feud.BaseBarony
}

var BBihac比哈奇 feud.Barony = &比哈奇BihacBarony{}

func init() {
	f := BBihac比哈奇.(*比哈奇BihacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bihac",
		TitleName: "比哈奇",
		TitleCode: "b_bihac",
	}
}
