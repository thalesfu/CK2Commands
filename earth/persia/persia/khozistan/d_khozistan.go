package khozistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/khozistan/avhaz"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/khozistan/khozistan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/khozistan/tigris"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhozistanDuke interface {
	feud.Duke
	CAvhaz阿瓦士() avhaz.AvhazCounty
	CKhozistan胡齐斯坦() khozistan.KhozistanCounty
	CTigris苏萨() tigris.TigrisCounty
}

type 胡齐斯坦KhozistanDuke struct {
	feud.BaseDuke
	阿瓦士Avhaz      avhaz.AvhazCounty
	胡齐斯坦Khozistan khozistan.KhozistanCounty
	苏萨Tigris      tigris.TigrisCounty
}

func (d *胡齐斯坦KhozistanDuke) CAvhaz阿瓦士() avhaz.AvhazCounty {
	return d.阿瓦士Avhaz
}

func (d *胡齐斯坦KhozistanDuke) CKhozistan胡齐斯坦() khozistan.KhozistanCounty {
	return d.胡齐斯坦Khozistan
}

func (d *胡齐斯坦KhozistanDuke) CTigris苏萨() tigris.TigrisCounty {
	return d.苏萨Tigris
}

var DKhozistan胡齐斯坦 KhozistanDuke = &胡齐斯坦KhozistanDuke{}

func init() {
	f := DKhozistan胡齐斯坦.(*胡齐斯坦KhozistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khozistan",
		TitleName: "胡齐斯坦",
		TitleCode: "d_khozistan",
		Counties:  map[string]feud.County{},
	}

	f.阿瓦士Avhaz = avhaz.CAvhaz阿瓦士
	f.阿瓦士Avhaz.SetParent(f)

	f.胡齐斯坦Khozistan = khozistan.CKhozistan胡齐斯坦
	f.胡齐斯坦Khozistan.SetParent(f)

	f.苏萨Tigris = tigris.CTigris苏萨
	f.苏萨Tigris.SetParent(f)

}
