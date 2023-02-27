package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基什塞本KisszebenBarony struct {
	feud.BaseBarony
}

var BKisszeben基什塞本 feud.Barony = &基什塞本KisszebenBarony{}

func init() {
    f := BKisszeben基什塞本.(*基什塞本KisszebenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kisszeben",
		TitleName: "基什塞本",
		TitleCode: "b_kisszeben",
	}
}
