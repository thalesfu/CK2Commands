package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 惠特比WhitbyBarony struct {
	feud.BaseBarony
}

var BWhitby惠特比 feud.Barony = &惠特比WhitbyBarony{}

func init() {
    f := BWhitby惠特比.(*惠特比WhitbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "whitby",
		TitleName: "惠特比",
		TitleCode: "b_whitby",
	}
}
