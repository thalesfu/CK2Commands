package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼姆法延NymphaionBarony struct {
	feud.BaseBarony
}

var BNymphaion尼姆法延 feud.Barony = &尼姆法延NymphaionBarony{}

func init() {
    f := BNymphaion尼姆法延.(*尼姆法延NymphaionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nymphaion",
		TitleName: "尼姆法延",
		TitleCode: "b_nymphaion",
	}
}
