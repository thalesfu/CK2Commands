package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔卡BirkaBarony struct {
	feud.BaseBarony
}

var BBirka比尔卡 feud.Barony = &比尔卡BirkaBarony{}

func init() {
	f := BBirka比尔卡.(*比尔卡BirkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birka",
		TitleName: "比尔卡",
		TitleCode: "b_birka",
	}
}
