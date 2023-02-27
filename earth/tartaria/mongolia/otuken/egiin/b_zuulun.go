package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卓伦ZuulunBarony struct {
	feud.BaseBarony
}

var BZuulun卓伦 feud.Barony = &卓伦ZuulunBarony{}

func init() {
    f := BZuulun卓伦.(*卓伦ZuulunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuulun",
		TitleName: "卓伦",
		TitleCode: "b_zuulun",
	}
}
