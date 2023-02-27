package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦克萨拉VaksalaBarony struct {
	feud.BaseBarony
}

var BVaksala瓦克萨拉 feud.Barony = &瓦克萨拉VaksalaBarony{}

func init() {
    f := BVaksala瓦克萨拉.(*瓦克萨拉VaksalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaksala",
		TitleName: "瓦克萨拉",
		TitleCode: "b_vaksala",
	}
}
