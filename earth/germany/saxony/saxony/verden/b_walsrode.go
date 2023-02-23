package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔斯罗德WalsrodeBarony struct {
	feud.BaseBarony
}

var BWalsrode瓦尔斯罗德 feud.Barony = &瓦尔斯罗德WalsrodeBarony{}

func init() {
	f := BWalsrode瓦尔斯罗德.(*瓦尔斯罗德WalsrodeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walsrode",
		TitleName: "瓦尔斯罗德",
		TitleCode: "b_walsrode",
	}
}
