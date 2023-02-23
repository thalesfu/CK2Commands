package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柏伦BurlunBarony struct {
	feud.BaseBarony
}

var BBurlun柏伦 feud.Barony = &柏伦BurlunBarony{}

func init() {
	f := BBurlun柏伦.(*柏伦BurlunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burlun",
		TitleName: "柏伦",
		TitleCode: "b_burlun",
	}
}
