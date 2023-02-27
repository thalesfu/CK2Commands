package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨乌达肯特SaudakentBarony struct {
	feud.BaseBarony
}

var BSaudakent萨乌达肯特 feud.Barony = &萨乌达肯特SaudakentBarony{}

func init() {
    f := BSaudakent萨乌达肯特.(*萨乌达肯特SaudakentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saudakent",
		TitleName: "萨乌达肯特",
		TitleCode: "b_saudakent",
	}
}
