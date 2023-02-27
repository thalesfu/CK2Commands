package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕克索斯PaxosBarony struct {
	feud.BaseBarony
}

var BPaxos帕克索斯 feud.Barony = &帕克索斯PaxosBarony{}

func init() {
    f := BPaxos帕克索斯.(*帕克索斯PaxosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paxos",
		TitleName: "帕克索斯",
		TitleCode: "b_paxos",
	}
}
