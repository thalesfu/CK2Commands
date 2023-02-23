package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 咩诃罗郁梨MehrauliBarony struct {
	feud.BaseBarony
}

var BMehrauli咩诃罗郁梨 feud.Barony = &咩诃罗郁梨MehrauliBarony{}

func init() {
	f := BMehrauli咩诃罗郁梨.(*咩诃罗郁梨MehrauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehrauli",
		TitleName: "咩诃罗郁梨",
		TitleCode: "b_mehrauli",
	}
}
