package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ToropetsCounty interface {
	feud.County
	BAndreapol安德烈亚波尔() feud.Barony
	BChistoye奇斯托耶() feud.Barony
	BDemyansk杰米扬斯克() feud.Barony
	BOkhvat奥赫瓦特() feud.Barony
	BPloskosh普洛斯科希() feud.Barony
	BToropets托罗佩茨() feud.Barony
	BZhizhitsa日日察() feud.Barony
}

type 托罗佩茨ToropetsCounty struct {
	feud.BaseCounty
	安德烈亚波尔Andreapol feud.Barony
	奇斯托耶Chistoye    feud.Barony
	杰米扬斯克Demyansk   feud.Barony
	奥赫瓦特Okhvat      feud.Barony
	普洛斯科希Ploskosh   feud.Barony
	托罗佩茨Toropets    feud.Barony
	日日察Zhizhitsa    feud.Barony
}

func (c *托罗佩茨ToropetsCounty) BAndreapol安德烈亚波尔() feud.Barony {
	return c.安德烈亚波尔Andreapol
}

func (c *托罗佩茨ToropetsCounty) BChistoye奇斯托耶() feud.Barony {
	return c.奇斯托耶Chistoye
}

func (c *托罗佩茨ToropetsCounty) BDemyansk杰米扬斯克() feud.Barony {
	return c.杰米扬斯克Demyansk
}

func (c *托罗佩茨ToropetsCounty) BOkhvat奥赫瓦特() feud.Barony {
	return c.奥赫瓦特Okhvat
}

func (c *托罗佩茨ToropetsCounty) BPloskosh普洛斯科希() feud.Barony {
	return c.普洛斯科希Ploskosh
}

func (c *托罗佩茨ToropetsCounty) BToropets托罗佩茨() feud.Barony {
	return c.托罗佩茨Toropets
}

func (c *托罗佩茨ToropetsCounty) BZhizhitsa日日察() feud.Barony {
	return c.日日察Zhizhitsa
}

var CToropets托罗佩茨 ToropetsCounty = &托罗佩茨ToropetsCounty{}

func init() {
	f := CToropets托罗佩茨.(*托罗佩茨ToropetsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "410",
		Title:     "toropets",
		TitleName: "托罗佩茨",
		TitleCode: "c_toropets",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德烈亚波尔Andreapol = BAndreapol安德烈亚波尔
	f.安德烈亚波尔Andreapol.SetParent(f)

	f.奇斯托耶Chistoye = BChistoye奇斯托耶
	f.奇斯托耶Chistoye.SetParent(f)

	f.杰米扬斯克Demyansk = BDemyansk杰米扬斯克
	f.杰米扬斯克Demyansk.SetParent(f)

	f.奥赫瓦特Okhvat = BOkhvat奥赫瓦特
	f.奥赫瓦特Okhvat.SetParent(f)

	f.普洛斯科希Ploskosh = BPloskosh普洛斯科希
	f.普洛斯科希Ploskosh.SetParent(f)

	f.托罗佩茨Toropets = BToropets托罗佩茨
	f.托罗佩茨Toropets.SetParent(f)

	f.日日察Zhizhitsa = BZhizhitsa日日察
	f.日日察Zhizhitsa.SetParent(f)

}
