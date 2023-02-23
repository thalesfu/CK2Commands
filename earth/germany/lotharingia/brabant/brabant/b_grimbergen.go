package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫林贝亨GrimbergenBarony struct {
	feud.BaseBarony
}

var BGrimbergen赫林贝亨 feud.Barony = &赫林贝亨GrimbergenBarony{}

func init() {
	f := BGrimbergen赫林贝亨.(*赫林贝亨GrimbergenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grimbergen",
		TitleName: "赫林贝亨",
		TitleCode: "b_grimbergen",
	}
}
