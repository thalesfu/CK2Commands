package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那婆陀NawadaBarony struct {
	feud.BaseBarony
}

var BNawada那婆陀 feud.Barony = &那婆陀NawadaBarony{}

func init() {
    f := BNawada那婆陀.(*那婆陀NawadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nawada",
		TitleName: "那婆陀",
		TitleCode: "b_nawada",
	}
}
