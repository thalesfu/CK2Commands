package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克奥吉AkaoudjBarony struct {
	feud.BaseBarony
}

var BAkaoudj阿克奥吉 feud.Barony = &阿克奥吉AkaoudjBarony{}

func init() {
    f := BAkaoudj阿克奥吉.(*阿克奥吉AkaoudjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akaoudj",
		TitleName: "阿克奥吉",
		TitleCode: "b_akaoudj",
	}
}
