package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩醯因陀罗城MahendranagarBarony struct {
	feud.BaseBarony
}

var BMahendranagar摩醯因陀罗城 feud.Barony = &摩醯因陀罗城MahendranagarBarony{}

func init() {
    f := BMahendranagar摩醯因陀罗城.(*摩醯因陀罗城MahendranagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahendranagar",
		TitleName: "摩醯因陀罗城",
		TitleCode: "b_mahendranagar",
	}
}
