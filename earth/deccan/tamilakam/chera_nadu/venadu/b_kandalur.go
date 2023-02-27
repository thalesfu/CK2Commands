package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘陀卢尔KandalurBarony struct {
	feud.BaseBarony
}

var BKandalur甘陀卢尔 feud.Barony = &甘陀卢尔KandalurBarony{}

func init() {
    f := BKandalur甘陀卢尔.(*甘陀卢尔KandalurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandalur",
		TitleName: "甘陀卢尔",
		TitleCode: "b_kandalur",
	}
}
