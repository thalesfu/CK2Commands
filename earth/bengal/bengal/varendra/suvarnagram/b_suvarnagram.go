package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏伐剌拿伽罗摩SuvarnagramBarony struct {
	feud.BaseBarony
}

var BSuvarnagram苏伐剌拿伽罗摩 feud.Barony = &苏伐剌拿伽罗摩SuvarnagramBarony{}

func init() {
	f := BSuvarnagram苏伐剌拿伽罗摩.(*苏伐剌拿伽罗摩SuvarnagramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suvarnagram",
		TitleName: "苏伐剌拿伽罗摩",
		TitleCode: "b_suvarnagram",
	}
}
