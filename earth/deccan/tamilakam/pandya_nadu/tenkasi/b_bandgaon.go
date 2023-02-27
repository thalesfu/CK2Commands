package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 槃陀伽罗摩BandgaonBarony struct {
	feud.BaseBarony
}

var BBandgaon槃陀伽罗摩 feud.Barony = &槃陀伽罗摩BandgaonBarony{}

func init() {
    f := BBandgaon槃陀伽罗摩.(*槃陀伽罗摩BandgaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandgaon",
		TitleName: "槃陀伽罗摩",
		TitleCode: "b_bandgaon",
	}
}
