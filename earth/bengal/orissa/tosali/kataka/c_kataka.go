package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KatakaCounty interface {
	feud.County
	BAthgarh八舍() feud.Barony
	BBhubaneswar部婆泥湿伐罗() feud.Barony
	BDhauli陶利() feud.Barony
	BKatak羯磔迦() feud.Barony
	BKonarak拘拿罗迦() feud.Barony
	BSakshigopal索崎瞿波罗() feud.Barony
	BSarala沙罗拉() feud.Barony
}

type 羯磔迦KatakaCounty struct {
	feud.BaseCounty
	八舍Athgarh         feud.Barony
	部婆泥湿伐罗Bhubaneswar feud.Barony
	陶利Dhauli          feud.Barony
	羯磔迦Katak          feud.Barony
	拘拿罗迦Konarak       feud.Barony
	索崎瞿波罗Sakshigopal  feud.Barony
	沙罗拉Sarala         feud.Barony
}

func (c *羯磔迦KatakaCounty) BAthgarh八舍() feud.Barony {
	return c.八舍Athgarh
}

func (c *羯磔迦KatakaCounty) BBhubaneswar部婆泥湿伐罗() feud.Barony {
	return c.部婆泥湿伐罗Bhubaneswar
}

func (c *羯磔迦KatakaCounty) BDhauli陶利() feud.Barony {
	return c.陶利Dhauli
}

func (c *羯磔迦KatakaCounty) BKatak羯磔迦() feud.Barony {
	return c.羯磔迦Katak
}

func (c *羯磔迦KatakaCounty) BKonarak拘拿罗迦() feud.Barony {
	return c.拘拿罗迦Konarak
}

func (c *羯磔迦KatakaCounty) BSakshigopal索崎瞿波罗() feud.Barony {
	return c.索崎瞿波罗Sakshigopal
}

func (c *羯磔迦KatakaCounty) BSarala沙罗拉() feud.Barony {
	return c.沙罗拉Sarala
}

var CKataka羯磔迦 KatakaCounty = &羯磔迦KatakaCounty{}

func init() {
	f := CKataka羯磔迦.(*羯磔迦KatakaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1129",
		Title:     "kataka",
		TitleName: "羯磔迦",
		TitleCode: "c_kataka",
		Baronies:  map[string]feud.Barony{},
	}

	f.八舍Athgarh = BAthgarh八舍
	f.八舍Athgarh.SetParent(f)

	f.部婆泥湿伐罗Bhubaneswar = BBhubaneswar部婆泥湿伐罗
	f.部婆泥湿伐罗Bhubaneswar.SetParent(f)

	f.陶利Dhauli = BDhauli陶利
	f.陶利Dhauli.SetParent(f)

	f.羯磔迦Katak = BKatak羯磔迦
	f.羯磔迦Katak.SetParent(f)

	f.拘拿罗迦Konarak = BKonarak拘拿罗迦
	f.拘拿罗迦Konarak.SetParent(f)

	f.索崎瞿波罗Sakshigopal = BSakshigopal索崎瞿波罗
	f.索崎瞿波罗Sakshigopal.SetParent(f)

	f.沙罗拉Sarala = BSarala沙罗拉
	f.沙罗拉Sarala.SetParent(f)

}
