package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArquesCounty interface {
	feud.County
	BArques阿尔克() feud.Barony
	BDieppe迪耶普() feud.Barony
	BFecamp费康() feud.Barony
	BHarfleur阿夫勒尔() feud.Barony
	BJumieges瑞米耶日() feud.Barony
	BLillebonne利勒博讷() feud.Barony
	BLongueville隆格维尔() feud.Barony
	BRouen鲁昂() feud.Barony
}

type 鲁昂ArquesCounty struct {
	feud.BaseCounty
	阿尔克Arques       feud.Barony
	迪耶普Dieppe       feud.Barony
	费康Fecamp        feud.Barony
	阿夫勒尔Harfleur    feud.Barony
	瑞米耶日Jumieges    feud.Barony
	利勒博讷Lillebonne  feud.Barony
	隆格维尔Longueville feud.Barony
	鲁昂Rouen         feud.Barony
}

func (c *鲁昂ArquesCounty) BArques阿尔克() feud.Barony {
	return c.阿尔克Arques
}

func (c *鲁昂ArquesCounty) BDieppe迪耶普() feud.Barony {
	return c.迪耶普Dieppe
}

func (c *鲁昂ArquesCounty) BFecamp费康() feud.Barony {
	return c.费康Fecamp
}

func (c *鲁昂ArquesCounty) BHarfleur阿夫勒尔() feud.Barony {
	return c.阿夫勒尔Harfleur
}

func (c *鲁昂ArquesCounty) BJumieges瑞米耶日() feud.Barony {
	return c.瑞米耶日Jumieges
}

func (c *鲁昂ArquesCounty) BLillebonne利勒博讷() feud.Barony {
	return c.利勒博讷Lillebonne
}

func (c *鲁昂ArquesCounty) BLongueville隆格维尔() feud.Barony {
	return c.隆格维尔Longueville
}

func (c *鲁昂ArquesCounty) BRouen鲁昂() feud.Barony {
	return c.鲁昂Rouen
}

var CArques鲁昂 ArquesCounty = &鲁昂ArquesCounty{}

func init() {
	f := CArques鲁昂.(*鲁昂ArquesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "97",
		Title:     "arques",
		TitleName: "鲁昂",
		TitleCode: "c_arques",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔克Arques = BArques阿尔克
	f.阿尔克Arques.SetParent(f)

	f.迪耶普Dieppe = BDieppe迪耶普
	f.迪耶普Dieppe.SetParent(f)

	f.费康Fecamp = BFecamp费康
	f.费康Fecamp.SetParent(f)

	f.阿夫勒尔Harfleur = BHarfleur阿夫勒尔
	f.阿夫勒尔Harfleur.SetParent(f)

	f.瑞米耶日Jumieges = BJumieges瑞米耶日
	f.瑞米耶日Jumieges.SetParent(f)

	f.利勒博讷Lillebonne = BLillebonne利勒博讷
	f.利勒博讷Lillebonne.SetParent(f)

	f.隆格维尔Longueville = BLongueville隆格维尔
	f.隆格维尔Longueville.SetParent(f)

	f.鲁昂Rouen = BRouen鲁昂
	f.鲁昂Rouen.SetParent(f)

}
