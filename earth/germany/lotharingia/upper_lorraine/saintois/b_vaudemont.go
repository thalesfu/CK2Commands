package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃代蒙VaudemontBarony struct {
	feud.BaseBarony
}

var BVaudemont沃代蒙 feud.Barony = &沃代蒙VaudemontBarony{}

func init() {
    f := BVaudemont沃代蒙.(*沃代蒙VaudemontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaudemont",
		TitleName: "沃代蒙",
		TitleCode: "b_vaudemont",
	}
}
