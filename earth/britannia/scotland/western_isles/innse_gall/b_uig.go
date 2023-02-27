package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌伊格UigBarony struct {
	feud.BaseBarony
}

var BUig乌伊格 feud.Barony = &乌伊格UigBarony{}

func init() {
    f := BUig乌伊格.(*乌伊格UigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uig",
		TitleName: "乌伊格",
		TitleCode: "b_uig",
	}
}
