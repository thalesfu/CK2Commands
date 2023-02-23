package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘迦提GamgadhiBarony struct {
	feud.BaseBarony
}

var BGamgadhi甘迦提 feud.Barony = &甘迦提GamgadhiBarony{}

func init() {
	f := BGamgadhi甘迦提.(*甘迦提GamgadhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gamgadhi",
		TitleName: "甘迦提",
		TitleCode: "b_gamgadhi",
	}
}
