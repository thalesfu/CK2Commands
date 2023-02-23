package scotland

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/albany"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/galloway"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/lothian"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/moray"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ScotlandKingdom interface {
	feud.Kingdom
	DAlbany奥尔巴尼() albany.AlbanyDuke
	DGalloway加洛韦() galloway.GallowayDuke
	DLothian洛锡安() lothian.LothianDuke
	DMoray马里() moray.MorayDuke
}

type 苏格兰ScotlandKingdom struct {
	feud.BaseKingdom
	奥尔巴尼Albany  albany.AlbanyDuke
	加洛韦Galloway galloway.GallowayDuke
	洛锡安Lothian  lothian.LothianDuke
	马里Moray     moray.MorayDuke
}

func (k *苏格兰ScotlandKingdom) DAlbany奥尔巴尼() albany.AlbanyDuke {
	return k.奥尔巴尼Albany
}

func (k *苏格兰ScotlandKingdom) DGalloway加洛韦() galloway.GallowayDuke {
	return k.加洛韦Galloway
}

func (k *苏格兰ScotlandKingdom) DLothian洛锡安() lothian.LothianDuke {
	return k.洛锡安Lothian
}

func (k *苏格兰ScotlandKingdom) DMoray马里() moray.MorayDuke {
	return k.马里Moray
}

var KScotland苏格兰 ScotlandKingdom = &苏格兰ScotlandKingdom{}

func init() {
	f := KScotland苏格兰.(*苏格兰ScotlandKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "scotland",
		TitleName: "苏格兰",
		TitleCode: "k_scotland",
		Dukes:     map[string]feud.Duke{},
	}

	f.奥尔巴尼Albany = albany.DAlbany奥尔巴尼
	f.奥尔巴尼Albany.SetParent(f)

	f.加洛韦Galloway = galloway.DGalloway加洛韦
	f.加洛韦Galloway.SetParent(f)

	f.洛锡安Lothian = lothian.DLothian洛锡安
	f.洛锡安Lothian.SetParent(f)

	f.马里Moray = moray.DMoray马里
	f.马里Moray.SetParent(f)

}
