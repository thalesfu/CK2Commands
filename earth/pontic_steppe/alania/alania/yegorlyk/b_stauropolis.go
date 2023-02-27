package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔夫罗波利斯StauropolisBarony struct {
	feud.BaseBarony
}

var BStauropolis斯塔夫罗波利斯 feud.Barony = &斯塔夫罗波利斯StauropolisBarony{}

func init() {
    f := BStauropolis斯塔夫罗波利斯.(*斯塔夫罗波利斯StauropolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stauropolis",
		TitleName: "斯塔夫罗波利斯",
		TitleCode: "b_stauropolis",
	}
}
