package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MontpellierCounty interface {
	feud.County
	BAiguesmortes艾格莫尔特() feud.Barony
	BBagnolssurceze塞兹河畔巴尼奥勒() feud.Barony
	BBeaucaire博凯尔() feud.Barony
	BMaguelone马格洛讷() feud.Barony
	BMelgueil莫吉奥() feud.Barony
	BMontpellier蒙彼利埃() feud.Barony
	BNimes尼姆() feud.Barony
	BSaintguilhemledesert圣吉扬_勒_代赛尔() feud.Barony
}

type 莫吉奥MontpellierCounty struct {
	feud.BaseCounty
	艾格莫尔特Aiguesmortes             feud.Barony
	塞兹河畔巴尼奥勒Bagnolssurceze        feud.Barony
	博凯尔Beaucaire                  feud.Barony
	马格洛讷Maguelone                 feud.Barony
	莫吉奥Melgueil                   feud.Barony
	蒙彼利埃Montpellier               feud.Barony
	尼姆Nimes                       feud.Barony
	圣吉扬_勒_代赛尔Saintguilhemledesert feud.Barony
}

func (c *莫吉奥MontpellierCounty) BAiguesmortes艾格莫尔特() feud.Barony {
	return c.艾格莫尔特Aiguesmortes
}

func (c *莫吉奥MontpellierCounty) BBagnolssurceze塞兹河畔巴尼奥勒() feud.Barony {
	return c.塞兹河畔巴尼奥勒Bagnolssurceze
}

func (c *莫吉奥MontpellierCounty) BBeaucaire博凯尔() feud.Barony {
	return c.博凯尔Beaucaire
}

func (c *莫吉奥MontpellierCounty) BMaguelone马格洛讷() feud.Barony {
	return c.马格洛讷Maguelone
}

func (c *莫吉奥MontpellierCounty) BMelgueil莫吉奥() feud.Barony {
	return c.莫吉奥Melgueil
}

func (c *莫吉奥MontpellierCounty) BMontpellier蒙彼利埃() feud.Barony {
	return c.蒙彼利埃Montpellier
}

func (c *莫吉奥MontpellierCounty) BNimes尼姆() feud.Barony {
	return c.尼姆Nimes
}

func (c *莫吉奥MontpellierCounty) BSaintguilhemledesert圣吉扬_勒_代赛尔() feud.Barony {
	return c.圣吉扬_勒_代赛尔Saintguilhemledesert
}

var CMontpellier莫吉奥 MontpellierCounty = &莫吉奥MontpellierCounty{}

func init() {
	f := CMontpellier莫吉奥.(*莫吉奥MontpellierCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "220",
		Title:     "montpellier",
		TitleName: "莫吉奥",
		TitleCode: "c_montpellier",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾格莫尔特Aiguesmortes = BAiguesmortes艾格莫尔特
	f.艾格莫尔特Aiguesmortes.SetParent(f)

	f.塞兹河畔巴尼奥勒Bagnolssurceze = BBagnolssurceze塞兹河畔巴尼奥勒
	f.塞兹河畔巴尼奥勒Bagnolssurceze.SetParent(f)

	f.博凯尔Beaucaire = BBeaucaire博凯尔
	f.博凯尔Beaucaire.SetParent(f)

	f.马格洛讷Maguelone = BMaguelone马格洛讷
	f.马格洛讷Maguelone.SetParent(f)

	f.莫吉奥Melgueil = BMelgueil莫吉奥
	f.莫吉奥Melgueil.SetParent(f)

	f.蒙彼利埃Montpellier = BMontpellier蒙彼利埃
	f.蒙彼利埃Montpellier.SetParent(f)

	f.尼姆Nimes = BNimes尼姆
	f.尼姆Nimes.SetParent(f)

	f.圣吉扬_勒_代赛尔Saintguilhemledesert = BSaintguilhemledesert圣吉扬_勒_代赛尔
	f.圣吉扬_勒_代赛尔Saintguilhemledesert.SetParent(f)

}
