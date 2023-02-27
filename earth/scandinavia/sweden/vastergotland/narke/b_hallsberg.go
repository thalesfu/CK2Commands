package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔斯贝里HallsbergBarony struct {
	feud.BaseBarony
}

var BHallsberg哈尔斯贝里 feud.Barony = &哈尔斯贝里HallsbergBarony{}

func init() {
    f := BHallsberg哈尔斯贝里.(*哈尔斯贝里HallsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hallsberg",
		TitleName: "哈尔斯贝里",
		TitleCode: "b_hallsberg",
	}
}
