package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗多补罗PithapuramBarony struct {
	feud.BaseBarony
}

var BPithapuram毗多补罗 feud.Barony = &毗多补罗PithapuramBarony{}

func init() {
    f := BPithapuram毗多补罗.(*毗多补罗PithapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pithapuram",
		TitleName: "毗多补罗",
		TitleCode: "b_pithapuram",
	}
}
