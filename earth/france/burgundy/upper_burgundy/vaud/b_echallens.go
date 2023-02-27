package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃沙朗EchallensBarony struct {
	feud.BaseBarony
}

var BEchallens埃沙朗 feud.Barony = &埃沙朗EchallensBarony{}

func init() {
    f := BEchallens埃沙朗.(*埃沙朗EchallensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "echallens",
		TitleName: "埃沙朗",
		TitleCode: "b_echallens",
	}
}
