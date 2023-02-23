package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托雷诺TorenoBarony struct {
	feud.BaseBarony
}

var BToreno托雷诺 feud.Barony = &托雷诺TorenoBarony{}

func init() {
	f := BToreno托雷诺.(*托雷诺TorenoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toreno",
		TitleName: "托雷诺",
		TitleCode: "b_toreno",
	}
}
