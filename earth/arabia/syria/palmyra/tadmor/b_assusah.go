package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏塞AssusahBarony struct {
	feud.BaseBarony
}

var BAssusah苏塞 feud.Barony = &苏塞AssusahBarony{}

func init() {
	f := BAssusah苏塞.(*苏塞AssusahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assusah",
		TitleName: "苏塞",
		TitleCode: "b_assusah",
	}
}
