package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小尖XiaojianBarony struct {
	feud.BaseBarony
}

var BXiaojian小尖 feud.Barony = &小尖XiaojianBarony{}

func init() {
    f := BXiaojian小尖.(*小尖XiaojianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiaojian",
		TitleName: "小尖",
		TitleCode: "b_xiaojian",
	}
}
