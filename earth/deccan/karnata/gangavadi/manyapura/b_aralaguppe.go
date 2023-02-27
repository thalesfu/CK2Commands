package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿剌罗加笈卑AralaguppeBarony struct {
	feud.BaseBarony
}

var BAralaguppe阿剌罗加笈卑 feud.Barony = &阿剌罗加笈卑AralaguppeBarony{}

func init() {
    f := BAralaguppe阿剌罗加笈卑.(*阿剌罗加笈卑AralaguppeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aralaguppe",
		TitleName: "阿剌罗加笈卑",
		TitleCode: "b_aralaguppe",
	}
}
