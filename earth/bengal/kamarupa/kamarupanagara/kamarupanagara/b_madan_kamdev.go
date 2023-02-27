package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩檀那迦摩提婆Madan_kamdevBarony struct {
	feud.BaseBarony
}

var BMadan_kamdev摩檀那迦摩提婆 feud.Barony = &摩檀那迦摩提婆Madan_kamdevBarony{}

func init() {
    f := BMadan_kamdev摩檀那迦摩提婆.(*摩檀那迦摩提婆Madan_kamdevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madan_kamdev",
		TitleName: "摩檀那迦摩提婆",
		TitleCode: "b_madan_kamdev",
	}
}
