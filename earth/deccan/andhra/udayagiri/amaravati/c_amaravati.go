package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmaravatiCounty interface {
	feud.County
	BAhobalam阿赫巴拉姆() feud.Barony
	BAmaravati阿摩罗伐底() feud.Barony
	BDhanyakataka驮那羯磔迦() feud.Barony
	BGurazala古罗扎拉() feud.Barony
	BKotappakonda拘吒波军荼() feud.Barony
	BMacherla马切尔拉() feud.Barony
	BVinukonda毗奴军荼() feud.Barony
}

type 阿摩罗伐底AmaravatiCounty struct {
	feud.BaseCounty
	阿赫巴拉姆Ahobalam     feud.Barony
	阿摩罗伐底Amaravati    feud.Barony
	驮那羯磔迦Dhanyakataka feud.Barony
	古罗扎拉Gurazala      feud.Barony
	拘吒波军荼Kotappakonda feud.Barony
	马切尔拉Macherla      feud.Barony
	毗奴军荼Vinukonda     feud.Barony
}

func (c *阿摩罗伐底AmaravatiCounty) BAhobalam阿赫巴拉姆() feud.Barony {
	return c.阿赫巴拉姆Ahobalam
}

func (c *阿摩罗伐底AmaravatiCounty) BAmaravati阿摩罗伐底() feud.Barony {
	return c.阿摩罗伐底Amaravati
}

func (c *阿摩罗伐底AmaravatiCounty) BDhanyakataka驮那羯磔迦() feud.Barony {
	return c.驮那羯磔迦Dhanyakataka
}

func (c *阿摩罗伐底AmaravatiCounty) BGurazala古罗扎拉() feud.Barony {
	return c.古罗扎拉Gurazala
}

func (c *阿摩罗伐底AmaravatiCounty) BKotappakonda拘吒波军荼() feud.Barony {
	return c.拘吒波军荼Kotappakonda
}

func (c *阿摩罗伐底AmaravatiCounty) BMacherla马切尔拉() feud.Barony {
	return c.马切尔拉Macherla
}

func (c *阿摩罗伐底AmaravatiCounty) BVinukonda毗奴军荼() feud.Barony {
	return c.毗奴军荼Vinukonda
}

var CAmaravati阿摩罗伐底 AmaravatiCounty = &阿摩罗伐底AmaravatiCounty{}

func init() {
	f := CAmaravati阿摩罗伐底.(*阿摩罗伐底AmaravatiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1207",
		Title:     "amaravati",
		TitleName: "阿摩罗伐底",
		TitleCode: "c_amaravati",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫巴拉姆Ahobalam = BAhobalam阿赫巴拉姆
	f.阿赫巴拉姆Ahobalam.SetParent(f)

	f.阿摩罗伐底Amaravati = BAmaravati阿摩罗伐底
	f.阿摩罗伐底Amaravati.SetParent(f)

	f.驮那羯磔迦Dhanyakataka = BDhanyakataka驮那羯磔迦
	f.驮那羯磔迦Dhanyakataka.SetParent(f)

	f.古罗扎拉Gurazala = BGurazala古罗扎拉
	f.古罗扎拉Gurazala.SetParent(f)

	f.拘吒波军荼Kotappakonda = BKotappakonda拘吒波军荼
	f.拘吒波军荼Kotappakonda.SetParent(f)

	f.马切尔拉Macherla = BMacherla马切尔拉
	f.马切尔拉Macherla.SetParent(f)

	f.毗奴军荼Vinukonda = BVinukonda毗奴军荼
	f.毗奴军荼Vinukonda.SetParent(f)

}
