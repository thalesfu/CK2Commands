package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴肯纳姆BuckenhamBarony struct {
	feud.BaseBarony
}

var BBuckenham巴肯纳姆 feud.Barony = &巴肯纳姆BuckenhamBarony{}

func init() {
    f := BBuckenham巴肯纳姆.(*巴肯纳姆BuckenhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buckenham",
		TitleName: "巴肯纳姆",
		TitleCode: "b_buckenham",
	}
}
