package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托罗ToroBarony struct {
	feud.BaseBarony
}

var BToro托罗 feud.Barony = &托罗ToroBarony{}

func init() {
	f := BToro托罗.(*托罗ToroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toro",
		TitleName: "托罗",
		TitleCode: "b_toro",
	}
}
