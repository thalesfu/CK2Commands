package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德拉斯坎AdraskanBarony struct {
	feud.BaseBarony
}

var BAdraskan阿德拉斯坎 feud.Barony = &阿德拉斯坎AdraskanBarony{}

func init() {
	f := BAdraskan阿德拉斯坎.(*阿德拉斯坎AdraskanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adraskan",
		TitleName: "阿德拉斯坎",
		TitleCode: "b_adraskan",
	}
}
