package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代赫里TehriBarony struct {
	feud.BaseBarony
}

var BTehri代赫里 feud.Barony = &代赫里TehriBarony{}

func init() {
	f := BTehri代赫里.(*代赫里TehriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tehri",
		TitleName: "代赫里",
		TitleCode: "b_tehri",
	}
}
