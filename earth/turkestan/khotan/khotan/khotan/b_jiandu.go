package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鞬都JianduBarony struct {
	feud.BaseBarony
}

var BJiandu鞬都 feud.Barony = &鞬都JianduBarony{}

func init() {
    f := BJiandu鞬都.(*鞬都JianduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiandu",
		TitleName: "鞬都",
		TitleCode: "b_jiandu",
	}
}
