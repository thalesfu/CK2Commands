package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托博尔斯克TobolskBarony struct {
	feud.BaseBarony
}

var BTobolsk托博尔斯克 feud.Barony = &托博尔斯克TobolskBarony{}

func init() {
    f := BTobolsk托博尔斯克.(*托博尔斯克TobolskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tobolsk",
		TitleName: "托博尔斯克",
		TitleCode: "b_tobolsk",
	}
}
