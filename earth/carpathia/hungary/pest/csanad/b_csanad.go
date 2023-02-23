package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔纳德CsanadBarony struct {
	feud.BaseBarony
}

var BCsanad乔纳德 feud.Barony = &乔纳德CsanadBarony{}

func init() {
	f := BCsanad乔纳德.(*乔纳德CsanadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csanad",
		TitleName: "乔纳德",
		TitleCode: "b_csanad",
	}
}
