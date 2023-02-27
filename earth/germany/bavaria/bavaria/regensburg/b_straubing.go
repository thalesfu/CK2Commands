package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施特劳宾StraubingBarony struct {
	feud.BaseBarony
}

var BStraubing施特劳宾 feud.Barony = &施特劳宾StraubingBarony{}

func init() {
    f := BStraubing施特劳宾.(*施特劳宾StraubingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "straubing",
		TitleName: "施特劳宾",
		TitleCode: "b_straubing",
	}
}
