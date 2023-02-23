package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥德瑟温GordoservonBarony struct {
	feud.BaseBarony
}

var BGordoservon哥德瑟温 feud.Barony = &哥德瑟温GordoservonBarony{}

func init() {
	f := BGordoservon哥德瑟温.(*哥德瑟温GordoservonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gordoservon",
		TitleName: "哥德瑟温",
		TitleCode: "b_gordoservon",
	}
}
