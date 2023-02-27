package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍洛夫基夫卡HolovkivkaBarony struct {
	feud.BaseBarony
}

var BHolovkivka霍洛夫基夫卡 feud.Barony = &霍洛夫基夫卡HolovkivkaBarony{}

func init() {
    f := BHolovkivka霍洛夫基夫卡.(*霍洛夫基夫卡HolovkivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holovkivka",
		TitleName: "霍洛夫基夫卡",
		TitleCode: "b_holovkivka",
	}
}
