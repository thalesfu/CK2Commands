package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝纳尔马德纳BenalmadenaBarony struct {
	feud.BaseBarony
}

var BBenalmadena贝纳尔马德纳 feud.Barony = &贝纳尔马德纳BenalmadenaBarony{}

func init() {
    f := BBenalmadena贝纳尔马德纳.(*贝纳尔马德纳BenalmadenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benalmadena",
		TitleName: "贝纳尔马德纳",
		TitleCode: "b_benalmadena",
	}
}
