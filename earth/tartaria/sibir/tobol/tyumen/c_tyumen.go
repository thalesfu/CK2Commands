package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyumenCounty interface {
	feud.County
	BBorovskiy博罗夫斯基() feud.Barony
	BNizhtavda下塔夫达() feud.Barony
	BNovtap诺夫塔普() feud.Barony
	BQashliq喀什里克() feud.Barony
	BSumkino苏姆基诺() feud.Barony
	BTobolsk托博尔斯克() feud.Barony
	BTugulym图古雷姆() feud.Barony
	BTyumen秋明() feud.Barony
}

type 成吉_图拉TyumenCounty struct {
	feud.BaseCounty
	博罗夫斯基Borovskiy feud.Barony
	下塔夫达Nizhtavda  feud.Barony
	诺夫塔普Novtap     feud.Barony
	喀什里克Qashliq    feud.Barony
	苏姆基诺Sumkino    feud.Barony
	托博尔斯克Tobolsk   feud.Barony
	图古雷姆Tugulym    feud.Barony
	秋明Tyumen       feud.Barony
}

func (c *成吉_图拉TyumenCounty) BBorovskiy博罗夫斯基() feud.Barony {
	return c.博罗夫斯基Borovskiy
}

func (c *成吉_图拉TyumenCounty) BNizhtavda下塔夫达() feud.Barony {
	return c.下塔夫达Nizhtavda
}

func (c *成吉_图拉TyumenCounty) BNovtap诺夫塔普() feud.Barony {
	return c.诺夫塔普Novtap
}

func (c *成吉_图拉TyumenCounty) BQashliq喀什里克() feud.Barony {
	return c.喀什里克Qashliq
}

func (c *成吉_图拉TyumenCounty) BSumkino苏姆基诺() feud.Barony {
	return c.苏姆基诺Sumkino
}

func (c *成吉_图拉TyumenCounty) BTobolsk托博尔斯克() feud.Barony {
	return c.托博尔斯克Tobolsk
}

func (c *成吉_图拉TyumenCounty) BTugulym图古雷姆() feud.Barony {
	return c.图古雷姆Tugulym
}

func (c *成吉_图拉TyumenCounty) BTyumen秋明() feud.Barony {
	return c.秋明Tyumen
}

var CTyumen成吉_图拉 TyumenCounty = &成吉_图拉TyumenCounty{}

func init() {
	f := CTyumen成吉_图拉.(*成吉_图拉TyumenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "891",
		Title:     "tyumen",
		TitleName: "成吉_图拉",
		TitleCode: "c_tyumen",
		Baronies:  map[string]feud.Barony{},
	}

	f.博罗夫斯基Borovskiy = BBorovskiy博罗夫斯基
	f.博罗夫斯基Borovskiy.SetParent(f)

	f.下塔夫达Nizhtavda = BNizhtavda下塔夫达
	f.下塔夫达Nizhtavda.SetParent(f)

	f.诺夫塔普Novtap = BNovtap诺夫塔普
	f.诺夫塔普Novtap.SetParent(f)

	f.喀什里克Qashliq = BQashliq喀什里克
	f.喀什里克Qashliq.SetParent(f)

	f.苏姆基诺Sumkino = BSumkino苏姆基诺
	f.苏姆基诺Sumkino.SetParent(f)

	f.托博尔斯克Tobolsk = BTobolsk托博尔斯克
	f.托博尔斯克Tobolsk.SetParent(f)

	f.图古雷姆Tugulym = BTugulym图古雷姆
	f.图古雷姆Tugulym.SetParent(f)

	f.秋明Tyumen = BTyumen秋明
	f.秋明Tyumen.SetParent(f)

}
