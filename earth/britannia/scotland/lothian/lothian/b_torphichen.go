package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托菲肯TorphichenBarony struct {
	feud.BaseBarony
}

var BTorphichen托菲肯 feud.Barony = &托菲肯TorphichenBarony{}

func init() {
    f := BTorphichen托菲肯.(*托菲肯TorphichenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torphichen",
		TitleName: "托菲肯",
		TitleCode: "b_torphichen",
	}
}
