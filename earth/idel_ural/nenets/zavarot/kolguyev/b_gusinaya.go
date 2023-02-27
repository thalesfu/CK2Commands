package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古西纳亚GusinayaBarony struct {
	feud.BaseBarony
}

var BGusinaya古西纳亚 feud.Barony = &古西纳亚GusinayaBarony{}

func init() {
    f := BGusinaya古西纳亚.(*古西纳亚GusinayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gusinaya",
		TitleName: "古西纳亚",
		TitleCode: "b_gusinaya",
	}
}
