package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅彻塔科比拉MechtakebiraBarony struct {
	feud.BaseBarony
}

var BMechtakebira梅彻塔科比拉 feud.Barony = &梅彻塔科比拉MechtakebiraBarony{}

func init() {
    f := BMechtakebira梅彻塔科比拉.(*梅彻塔科比拉MechtakebiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mechtakebira",
		TitleName: "梅彻塔科比拉",
		TitleCode: "b_mechtakebira",
	}
}
