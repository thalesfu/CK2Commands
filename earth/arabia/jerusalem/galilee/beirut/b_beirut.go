package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝鲁特BeirutBarony struct {
	feud.BaseBarony
}

var BBeirut贝鲁特 feud.Barony = &贝鲁特BeirutBarony{}

func init() {
	f := BBeirut贝鲁特.(*贝鲁特BeirutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beirut",
		TitleName: "贝鲁特",
		TitleCode: "b_beirut",
	}
}
