package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏塞港MarsasusahBarony struct {
	feud.BaseBarony
}

var BMarsasusah苏塞港 feud.Barony = &苏塞港MarsasusahBarony{}

func init() {
	f := BMarsasusah苏塞港.(*苏塞港MarsasusahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marsasusah",
		TitleName: "苏塞港",
		TitleCode: "b_marsasusah",
	}
}
