package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠罗伽罗VairagaraBarony struct {
	feud.BaseBarony
}

var BVairagara吠罗伽罗 feud.Barony = &吠罗伽罗VairagaraBarony{}

func init() {
	f := BVairagara吠罗伽罗.(*吠罗伽罗VairagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vairagara",
		TitleName: "吠罗伽罗",
		TitleCode: "b_vairagara",
	}
}
