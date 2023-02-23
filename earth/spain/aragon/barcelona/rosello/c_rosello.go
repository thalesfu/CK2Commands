package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RoselloCounty interface {
	feud.County
	BCanigo卡尼戈() feud.Barony
	BCeret塞雷特() feud.Barony
	BCotlliure科利乌尔() feud.Barony
	BCuixa库沙() feud.Barony
	BElna埃尔讷() feud.Barony
	BOltrera奥尔特雷拉() feud.Barony
	BPerpinya佩皮尼昂() feud.Barony
	BPrada普拉达() feud.Barony
}

type 罗塞约RoselloCounty struct {
	feud.BaseCounty
	卡尼戈Canigo     feud.Barony
	塞雷特Ceret      feud.Barony
	科利乌尔Cotlliure feud.Barony
	库沙Cuixa       feud.Barony
	埃尔讷Elna       feud.Barony
	奥尔特雷拉Oltrera  feud.Barony
	佩皮尼昂Perpinya  feud.Barony
	普拉达Prada      feud.Barony
}

func (c *罗塞约RoselloCounty) BCanigo卡尼戈() feud.Barony {
	return c.卡尼戈Canigo
}

func (c *罗塞约RoselloCounty) BCeret塞雷特() feud.Barony {
	return c.塞雷特Ceret
}

func (c *罗塞约RoselloCounty) BCotlliure科利乌尔() feud.Barony {
	return c.科利乌尔Cotlliure
}

func (c *罗塞约RoselloCounty) BCuixa库沙() feud.Barony {
	return c.库沙Cuixa
}

func (c *罗塞约RoselloCounty) BElna埃尔讷() feud.Barony {
	return c.埃尔讷Elna
}

func (c *罗塞约RoselloCounty) BOltrera奥尔特雷拉() feud.Barony {
	return c.奥尔特雷拉Oltrera
}

func (c *罗塞约RoselloCounty) BPerpinya佩皮尼昂() feud.Barony {
	return c.佩皮尼昂Perpinya
}

func (c *罗塞约RoselloCounty) BPrada普拉达() feud.Barony {
	return c.普拉达Prada
}

var CRosello罗塞约 RoselloCounty = &罗塞约RoselloCounty{}

func init() {
	f := CRosello罗塞约.(*罗塞约RoselloCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "211",
		Title:     "rosello",
		TitleName: "罗塞约",
		TitleCode: "c_rosello",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡尼戈Canigo = BCanigo卡尼戈
	f.卡尼戈Canigo.SetParent(f)

	f.塞雷特Ceret = BCeret塞雷特
	f.塞雷特Ceret.SetParent(f)

	f.科利乌尔Cotlliure = BCotlliure科利乌尔
	f.科利乌尔Cotlliure.SetParent(f)

	f.库沙Cuixa = BCuixa库沙
	f.库沙Cuixa.SetParent(f)

	f.埃尔讷Elna = BElna埃尔讷
	f.埃尔讷Elna.SetParent(f)

	f.奥尔特雷拉Oltrera = BOltrera奥尔特雷拉
	f.奥尔特雷拉Oltrera.SetParent(f)

	f.佩皮尼昂Perpinya = BPerpinya佩皮尼昂
	f.佩皮尼昂Perpinya.SetParent(f)

	f.普拉达Prada = BPrada普拉达
	f.普拉达Prada.SetParent(f)

}
