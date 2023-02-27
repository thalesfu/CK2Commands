package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗罗知湿伐罗VirateshwarBarony struct {
	feud.BaseBarony
}

var BVirateshwar毗罗知湿伐罗 feud.Barony = &毗罗知湿伐罗VirateshwarBarony{}

func init() {
    f := BVirateshwar毗罗知湿伐罗.(*毗罗知湿伐罗VirateshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "virateshwar",
		TitleName: "毗罗知湿伐罗",
		TitleCode: "b_virateshwar",
	}
}
