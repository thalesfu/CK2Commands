package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博季努尔德BojnurdBarony struct {
	feud.BaseBarony
}

var BBojnurd博季努尔德 feud.Barony = &博季努尔德BojnurdBarony{}

func init() {
    f := BBojnurd博季努尔德.(*博季努尔德BojnurdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bojnurd",
		TitleName: "博季努尔德",
		TitleCode: "b_bojnurd",
	}
}
