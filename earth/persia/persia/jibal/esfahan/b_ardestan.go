package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔德斯坦ArdestanBarony struct {
	feud.BaseBarony
}

var BArdestan阿尔德斯坦 feud.Barony = &阿尔德斯坦ArdestanBarony{}

func init() {
    f := BArdestan阿尔德斯坦.(*阿尔德斯坦ArdestanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardestan",
		TitleName: "阿尔德斯坦",
		TitleCode: "b_ardestan",
	}
}
