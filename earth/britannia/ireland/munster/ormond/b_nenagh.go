package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼纳NenaghBarony struct {
	feud.BaseBarony
}

var BNenagh尼纳 feud.Barony = &尼纳NenaghBarony{}

func init() {
    f := BNenagh尼纳.(*尼纳NenaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nenagh",
		TitleName: "尼纳",
		TitleCode: "b_nenagh",
	}
}
