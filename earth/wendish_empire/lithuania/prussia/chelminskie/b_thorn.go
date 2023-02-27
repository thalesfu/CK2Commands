package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托恩ThornBarony struct {
	feud.BaseBarony
}

var BThorn托恩 feud.Barony = &托恩ThornBarony{}

func init() {
    f := BThorn托恩.(*托恩ThornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thorn",
		TitleName: "托恩",
		TitleCode: "b_thorn",
	}
}
