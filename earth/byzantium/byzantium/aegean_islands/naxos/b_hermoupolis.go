package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔穆波利HermoupolisBarony struct {
	feud.BaseBarony
}

var BHermoupolis埃尔穆波利 feud.Barony = &埃尔穆波利HermoupolisBarony{}

func init() {
    f := BHermoupolis埃尔穆波利.(*埃尔穆波利HermoupolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hermoupolis",
		TitleName: "埃尔穆波利",
		TitleCode: "b_hermoupolis",
	}
}
