package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩尼迦突伽ManikdurgBarony struct {
	feud.BaseBarony
}

var BManikdurg摩尼迦突伽 feud.Barony = &摩尼迦突伽ManikdurgBarony{}

func init() {
    f := BManikdurg摩尼迦突伽.(*摩尼迦突伽ManikdurgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manikdurg",
		TitleName: "摩尼迦突伽",
		TitleCode: "b_manikdurg",
	}
}
