package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳克索斯NaxosBarony struct {
	feud.BaseBarony
}

var BNaxos纳克索斯 feud.Barony = &纳克索斯NaxosBarony{}

func init() {
    f := BNaxos纳克索斯.(*纳克索斯NaxosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naxos",
		TitleName: "纳克索斯",
		TitleCode: "b_naxos",
	}
}
