package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉姆西ClamecyBarony struct {
	feud.BaseBarony
}

var BClamecy克拉姆西 feud.Barony = &克拉姆西ClamecyBarony{}

func init() {
    f := BClamecy克拉姆西.(*克拉姆西ClamecyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clamecy",
		TitleName: "克拉姆西",
		TitleCode: "b_clamecy",
	}
}
