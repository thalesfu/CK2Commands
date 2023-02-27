package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克茅斯CockermouthBarony struct {
	feud.BaseBarony
}

var BCockermouth科克茅斯 feud.Barony = &科克茅斯CockermouthBarony{}

func init() {
    f := BCockermouth科克茅斯.(*科克茅斯CockermouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cockermouth",
		TitleName: "科克茅斯",
		TitleCode: "b_cockermouth",
	}
}
