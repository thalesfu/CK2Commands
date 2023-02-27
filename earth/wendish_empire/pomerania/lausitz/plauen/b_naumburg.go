package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙姆堡NaumburgBarony struct {
	feud.BaseBarony
}

var BNaumburg瑙姆堡 feud.Barony = &瑙姆堡NaumburgBarony{}

func init() {
    f := BNaumburg瑙姆堡.(*瑙姆堡NaumburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naumburg",
		TitleName: "瑙姆堡",
		TitleCode: "b_naumburg",
	}
}
