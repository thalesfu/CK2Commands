package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦根费尔德WagenfeldBarony struct {
	feud.BaseBarony
}

var BWagenfeld瓦根费尔德 feud.Barony = &瓦根费尔德WagenfeldBarony{}

func init() {
    f := BWagenfeld瓦根费尔德.(*瓦根费尔德WagenfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wagenfeld",
		TitleName: "瓦根费尔德",
		TitleCode: "b_wagenfeld",
	}
}
