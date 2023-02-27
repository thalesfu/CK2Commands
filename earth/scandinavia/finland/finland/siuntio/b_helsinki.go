package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔辛基HelsinkiBarony struct {
	feud.BaseBarony
}

var BHelsinki赫尔辛基 feud.Barony = &赫尔辛基HelsinkiBarony{}

func init() {
    f := BHelsinki赫尔辛基.(*赫尔辛基HelsinkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helsinki",
		TitleName: "赫尔辛基",
		TitleCode: "b_helsinki",
	}
}
