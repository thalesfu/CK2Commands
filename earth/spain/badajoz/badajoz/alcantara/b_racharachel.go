package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉查拉切尔RacharachelBarony struct {
	feud.BaseBarony
}

var BRacharachel拉查拉切尔 feud.Barony = &拉查拉切尔RacharachelBarony{}

func init() {
    f := BRacharachel拉查拉切尔.(*拉查拉切尔RacharachelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "racharachel",
		TitleName: "拉查拉切尔",
		TitleCode: "b_racharachel",
	}
}
