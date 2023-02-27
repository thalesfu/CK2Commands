package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊瑟尔堡IsselburgBarony struct {
	feud.BaseBarony
}

var BIsselburg伊瑟尔堡 feud.Barony = &伊瑟尔堡IsselburgBarony{}

func init() {
    f := BIsselburg伊瑟尔堡.(*伊瑟尔堡IsselburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isselburg",
		TitleName: "伊瑟尔堡",
		TitleCode: "b_isselburg",
	}
}
