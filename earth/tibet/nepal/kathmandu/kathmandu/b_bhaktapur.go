package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 薄多补罗BhaktapurBarony struct {
	feud.BaseBarony
}

var BBhaktapur薄多补罗 feud.Barony = &薄多补罗BhaktapurBarony{}

func init() {
    f := BBhaktapur薄多补罗.(*薄多补罗BhaktapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhaktapur",
		TitleName: "薄多补罗",
		TitleCode: "b_bhaktapur",
	}
}
