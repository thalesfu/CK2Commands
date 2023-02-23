package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔德胡斯VardohusBarony struct {
	feud.BaseBarony
}

var BVardohus瓦尔德胡斯 feud.Barony = &瓦尔德胡斯VardohusBarony{}

func init() {
	f := BVardohus瓦尔德胡斯.(*瓦尔德胡斯VardohusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vardohus",
		TitleName: "瓦尔德胡斯",
		TitleCode: "b_vardohus",
	}
}
