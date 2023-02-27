package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔卡亚特WolqaytBarony struct {
	feud.BaseBarony
}

var BWolqayt瓦尔卡亚特 feud.Barony = &瓦尔卡亚特WolqaytBarony{}

func init() {
    f := BWolqayt瓦尔卡亚特.(*瓦尔卡亚特WolqaytBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolqayt",
		TitleName: "瓦尔卡亚特",
		TitleCode: "b_wolqayt",
	}
}
