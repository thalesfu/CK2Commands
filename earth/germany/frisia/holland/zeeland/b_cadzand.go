package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡德赞德CadzandBarony struct {
	feud.BaseBarony
}

var BCadzand卡德赞德 feud.Barony = &卡德赞德CadzandBarony{}

func init() {
    f := BCadzand卡德赞德.(*卡德赞德CadzandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cadzand",
		TitleName: "卡德赞德",
		TitleCode: "b_cadzand",
	}
}
