package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 若瑟兰JosselinBarony struct {
	feud.BaseBarony
}

var BJosselin若瑟兰 feud.Barony = &若瑟兰JosselinBarony{}

func init() {
    f := BJosselin若瑟兰.(*若瑟兰JosselinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "josselin",
		TitleName: "若瑟兰",
		TitleCode: "b_josselin",
	}
}
