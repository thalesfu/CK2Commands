package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悉多补罗SitpurBarony struct {
	feud.BaseBarony
}

var BSitpur悉多补罗 feud.Barony = &悉多补罗SitpurBarony{}

func init() {
    f := BSitpur悉多补罗.(*悉多补罗SitpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sitpur",
		TitleName: "悉多补罗",
		TitleCode: "b_sitpur",
	}
}
