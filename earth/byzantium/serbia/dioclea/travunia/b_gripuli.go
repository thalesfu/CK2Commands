package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格里普利GripuliBarony struct {
	feud.BaseBarony
}

var BGripuli格里普利 feud.Barony = &格里普利GripuliBarony{}

func init() {
	f := BGripuli格里普利.(*格里普利GripuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gripuli",
		TitleName: "格里普利",
		TitleCode: "b_gripuli",
	}
}
