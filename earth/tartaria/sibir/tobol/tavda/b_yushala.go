package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤沙拉YushalaBarony struct {
	feud.BaseBarony
}

var BYushala尤沙拉 feud.Barony = &尤沙拉YushalaBarony{}

func init() {
    f := BYushala尤沙拉.(*尤沙拉YushalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yushala",
		TitleName: "尤沙拉",
		TitleCode: "b_yushala",
	}
}
