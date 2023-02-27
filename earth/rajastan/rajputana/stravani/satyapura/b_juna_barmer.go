package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周那婆摩尔Juna_barmerBarony struct {
	feud.BaseBarony
}

var BJuna_barmer周那婆摩尔 feud.Barony = &周那婆摩尔Juna_barmerBarony{}

func init() {
    f := BJuna_barmer周那婆摩尔.(*周那婆摩尔Juna_barmerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juna_barmer",
		TitleName: "周那婆摩尔",
		TitleCode: "b_juna_barmer",
	}
}
