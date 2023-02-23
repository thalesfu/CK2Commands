package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿什卡莱AskaleBarony struct {
	feud.BaseBarony
}

var BAskale阿什卡莱 feud.Barony = &阿什卡莱AskaleBarony{}

func init() {
	f := BAskale阿什卡莱.(*阿什卡莱AskaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askale",
		TitleName: "阿什卡莱",
		TitleCode: "b_askale",
	}
}
