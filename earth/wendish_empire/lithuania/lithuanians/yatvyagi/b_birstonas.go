package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔什托纳斯BirstonasBarony struct {
	feud.BaseBarony
}

var BBirstonas比尔什托纳斯 feud.Barony = &比尔什托纳斯BirstonasBarony{}

func init() {
    f := BBirstonas比尔什托纳斯.(*比尔什托纳斯BirstonasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birstonas",
		TitleName: "比尔什托纳斯",
		TitleCode: "b_birstonas",
	}
}
