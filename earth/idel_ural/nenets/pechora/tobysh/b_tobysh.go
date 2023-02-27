package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托贝什TobyshBarony struct {
	feud.BaseBarony
}

var BTobysh托贝什 feud.Barony = &托贝什TobyshBarony{}

func init() {
    f := BTobysh托贝什.(*托贝什TobyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tobysh",
		TitleName: "托贝什",
		TitleCode: "b_tobysh",
	}
}
