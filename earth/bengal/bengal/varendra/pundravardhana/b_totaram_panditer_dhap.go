package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豆多蓝般提蒂罗土丘Totaram_panditer_dhapBarony struct {
	feud.BaseBarony
}

var BTotaram_panditer_dhap豆多蓝般提蒂罗土丘 feud.Barony = &豆多蓝般提蒂罗土丘Totaram_panditer_dhapBarony{}

func init() {
    f := BTotaram_panditer_dhap豆多蓝般提蒂罗土丘.(*豆多蓝般提蒂罗土丘Totaram_panditer_dhapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "totaram_panditer_dhap",
		TitleName: "豆多蓝般提蒂罗土丘",
		TitleCode: "b_totaram_panditer_dhap",
	}
}
