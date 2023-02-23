package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦罗埃MeruwahBarony struct {
	feud.BaseBarony
}

var BMeruwah麦罗埃 feud.Barony = &麦罗埃MeruwahBarony{}

func init() {
	f := BMeruwah麦罗埃.(*麦罗埃MeruwahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meruwah",
		TitleName: "麦罗埃",
		TitleCode: "b_meruwah",
	}
}
