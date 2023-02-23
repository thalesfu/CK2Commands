package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆塔帕里MutapaliBarony struct {
	feud.BaseBarony
}

var BMutapali穆塔帕里 feud.Barony = &穆塔帕里MutapaliBarony{}

func init() {
	f := BMutapali穆塔帕里.(*穆塔帕里MutapaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutapali",
		TitleName: "穆塔帕里",
		TitleCode: "b_mutapali",
	}
}
