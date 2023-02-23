package herat

import (
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/herat/badghis"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/herat/herat"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/herat/mandesh"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HeratDuke interface {
	feud.Duke
	CBadghis巴德吉斯() badghis.BadghisCounty
	CHerat亦鲁() herat.HeratCounty
	CMandesh曼德什() mandesh.MandeshCounty
}

type 亦鲁HeratDuke struct {
	feud.BaseDuke
	巴德吉斯Badghis badghis.BadghisCounty
	亦鲁Herat     herat.HeratCounty
	曼德什Mandesh  mandesh.MandeshCounty
}

func (d *亦鲁HeratDuke) CBadghis巴德吉斯() badghis.BadghisCounty {
	return d.巴德吉斯Badghis
}

func (d *亦鲁HeratDuke) CHerat亦鲁() herat.HeratCounty {
	return d.亦鲁Herat
}

func (d *亦鲁HeratDuke) CMandesh曼德什() mandesh.MandeshCounty {
	return d.曼德什Mandesh
}

var DHerat亦鲁 HeratDuke = &亦鲁HeratDuke{}

func init() {
	f := DHerat亦鲁.(*亦鲁HeratDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "herat",
		TitleName: "亦鲁",
		TitleCode: "d_herat",
		Counties:  map[string]feud.County{},
	}

	f.巴德吉斯Badghis = badghis.CBadghis巴德吉斯
	f.巴德吉斯Badghis.SetParent(f)

	f.亦鲁Herat = herat.CHerat亦鲁
	f.亦鲁Herat.SetParent(f)

	f.曼德什Mandesh = mandesh.CMandesh曼德什
	f.曼德什Mandesh.SetParent(f)

}
