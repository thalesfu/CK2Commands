package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤甘YuganBarony struct {
	feud.BaseBarony
}

var BYugan尤甘 feud.Barony = &尤甘YuganBarony{}

func init() {
    f := BYugan尤甘.(*尤甘YuganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yugan",
		TitleName: "尤甘",
		TitleCode: "b_yugan",
	}
}
