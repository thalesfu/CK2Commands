package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼荼罗姞利呬MandalgarhBarony struct {
	feud.BaseBarony
}

var BMandalgarh曼荼罗姞利呬 feud.Barony = &曼荼罗姞利呬MandalgarhBarony{}

func init() {
    f := BMandalgarh曼荼罗姞利呬.(*曼荼罗姞利呬MandalgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandalgarh",
		TitleName: "曼荼罗姞利呬",
		TitleCode: "b_mandalgarh",
	}
}
