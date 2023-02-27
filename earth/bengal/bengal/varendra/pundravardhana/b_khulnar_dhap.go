package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘罗土丘Khulnar_dhapBarony struct {
	feud.BaseBarony
}

var BKhulnar_dhap丘罗土丘 feud.Barony = &丘罗土丘Khulnar_dhapBarony{}

func init() {
    f := BKhulnar_dhap丘罗土丘.(*丘罗土丘Khulnar_dhapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khulnar_dhap",
		TitleName: "丘罗土丘",
		TitleCode: "b_khulnar_dhap",
	}
}
