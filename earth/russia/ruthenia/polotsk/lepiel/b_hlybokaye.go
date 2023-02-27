package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格卢博科耶HlybokayeBarony struct {
	feud.BaseBarony
}

var BHlybokaye格卢博科耶 feud.Barony = &格卢博科耶HlybokayeBarony{}

func init() {
    f := BHlybokaye格卢博科耶.(*格卢博科耶HlybokayeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hlybokaye",
		TitleName: "格卢博科耶",
		TitleCode: "b_hlybokaye",
	}
}
