package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃陀瞿荼BandagudaBarony struct {
	feud.BaseBarony
}

var BBandaguda槃陀瞿荼 feud.Barony = &槃陀瞿荼BandagudaBarony{}

func init() {
	f := BBandaguda槃陀瞿荼.(*槃陀瞿荼BandagudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandaguda",
		TitleName: "槃陀瞿荼",
		TitleCode: "b_bandaguda",
	}
}
