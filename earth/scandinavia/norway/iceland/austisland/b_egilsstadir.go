package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃伊尔斯塔济EgilsstadirBarony struct {
	feud.BaseBarony
}

var BEgilsstadir埃伊尔斯塔济 feud.Barony = &埃伊尔斯塔济EgilsstadirBarony{}

func init() {
	f := BEgilsstadir埃伊尔斯塔济.(*埃伊尔斯塔济EgilsstadirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egilsstadir",
		TitleName: "埃伊尔斯塔济",
		TitleCode: "b_egilsstadir",
	}
}
