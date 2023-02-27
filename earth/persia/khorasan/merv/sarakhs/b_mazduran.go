package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马兹杜兰MazduranBarony struct {
	feud.BaseBarony
}

var BMazduran马兹杜兰 feud.Barony = &马兹杜兰MazduranBarony{}

func init() {
    f := BMazduran马兹杜兰.(*马兹杜兰MazduranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazduran",
		TitleName: "马兹杜兰",
		TitleCode: "b_mazduran",
	}
}
