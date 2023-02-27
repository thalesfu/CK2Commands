package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶洛沃YelovoBarony struct {
	feud.BaseBarony
}

var BYelovo叶洛沃 feud.Barony = &叶洛沃YelovoBarony{}

func init() {
    f := BYelovo叶洛沃.(*叶洛沃YelovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yelovo",
		TitleName: "叶洛沃",
		TitleCode: "b_yelovo",
	}
}
