package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔卡ZarqaBarony struct {
	feud.BaseBarony
}

var BZarqa扎尔卡 feud.Barony = &扎尔卡ZarqaBarony{}

func init() {
    f := BZarqa扎尔卡.(*扎尔卡ZarqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarqa",
		TitleName: "扎尔卡",
		TitleCode: "b_zarqa",
	}
}
