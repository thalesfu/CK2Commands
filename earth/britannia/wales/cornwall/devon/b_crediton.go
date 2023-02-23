package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷迪顿CreditonBarony struct {
	feud.BaseBarony
}

var BCrediton克雷迪顿 feud.Barony = &克雷迪顿CreditonBarony{}

func init() {
	f := BCrediton克雷迪顿.(*克雷迪顿CreditonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crediton",
		TitleName: "克雷迪顿",
		TitleCode: "b_crediton",
	}
}
