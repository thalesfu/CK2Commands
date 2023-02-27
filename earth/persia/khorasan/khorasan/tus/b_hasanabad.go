package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈桑阿巴德HasanabadBarony struct {
	feud.BaseBarony
}

var BHasanabad哈桑阿巴德 feud.Barony = &哈桑阿巴德HasanabadBarony{}

func init() {
    f := BHasanabad哈桑阿巴德.(*哈桑阿巴德HasanabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hasanabad",
		TitleName: "哈桑阿巴德",
		TitleCode: "b_hasanabad",
	}
}
