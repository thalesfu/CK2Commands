package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海尔斯堡HeilsbergBarony struct {
	feud.BaseBarony
}

var BHeilsberg海尔斯堡 feud.Barony = &海尔斯堡HeilsbergBarony{}

func init() {
    f := BHeilsberg海尔斯堡.(*海尔斯堡HeilsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heilsberg",
		TitleName: "海尔斯堡",
		TitleCode: "b_heilsberg",
	}
}
