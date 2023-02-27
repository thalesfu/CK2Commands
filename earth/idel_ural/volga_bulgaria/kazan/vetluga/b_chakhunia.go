package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察胡尼亚ChakhuniaBarony struct {
	feud.BaseBarony
}

var BChakhunia察胡尼亚 feud.Barony = &察胡尼亚ChakhuniaBarony{}

func init() {
    f := BChakhunia察胡尼亚.(*察胡尼亚ChakhuniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chakhunia",
		TitleName: "察胡尼亚",
		TitleCode: "b_chakhunia",
	}
}
