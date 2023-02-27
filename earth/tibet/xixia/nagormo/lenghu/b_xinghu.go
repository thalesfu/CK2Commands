package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 星湖XinghuBarony struct {
	feud.BaseBarony
}

var BXinghu星湖 feud.Barony = &星湖XinghuBarony{}

func init() {
    f := BXinghu星湖.(*星湖XinghuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xinghu",
		TitleName: "星湖",
		TitleCode: "b_xinghu",
	}
}
