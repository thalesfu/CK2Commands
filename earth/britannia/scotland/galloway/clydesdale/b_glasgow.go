package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉斯哥GlasgowBarony struct {
	feud.BaseBarony
}

var BGlasgow格拉斯哥 feud.Barony = &格拉斯哥GlasgowBarony{}

func init() {
    f := BGlasgow格拉斯哥.(*格拉斯哥GlasgowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glasgow",
		TitleName: "格拉斯哥",
		TitleCode: "b_glasgow",
	}
}
