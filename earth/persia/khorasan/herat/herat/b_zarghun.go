package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔贡ZarghunBarony struct {
	feud.BaseBarony
}

var BZarghun扎尔贡 feud.Barony = &扎尔贡ZarghunBarony{}

func init() {
    f := BZarghun扎尔贡.(*扎尔贡ZarghunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarghun",
		TitleName: "扎尔贡",
		TitleCode: "b_zarghun",
	}
}
