package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托斯诺TosnoBarony struct {
	feud.BaseBarony
}

var BTosno托斯诺 feud.Barony = &托斯诺TosnoBarony{}

func init() {
    f := BTosno托斯诺.(*托斯诺TosnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tosno",
		TitleName: "托斯诺",
		TitleCode: "b_tosno",
	}
}
