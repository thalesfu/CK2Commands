package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺切托NocetoBarony struct {
	feud.BaseBarony
}

var BNoceto诺切托 feud.Barony = &诺切托NocetoBarony{}

func init() {
	f := BNoceto诺切托.(*诺切托NocetoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noceto",
		TitleName: "诺切托",
		TitleCode: "b_noceto",
	}
}
