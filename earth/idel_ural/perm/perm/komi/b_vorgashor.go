package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔加绍尔VorgashorBarony struct {
	feud.BaseBarony
}

var BVorgashor沃尔加绍尔 feud.Barony = &沃尔加绍尔VorgashorBarony{}

func init() {
    f := BVorgashor沃尔加绍尔.(*沃尔加绍尔VorgashorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vorgashor",
		TitleName: "沃尔加绍尔",
		TitleCode: "b_vorgashor",
	}
}
