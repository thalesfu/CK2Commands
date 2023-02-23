package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmiensCounty interface {
	feud.County
	BAmiens亚眠() feud.Barony
	BBreteuil布勒特伊() feud.Barony
	BCorbie科尔比() feud.Barony
	BMontdidier蒙迪迪耶() feud.Barony
	BNesle内勒() feud.Barony
	BNoyon努瓦永() feud.Barony
	BPeronne佩罗讷() feud.Barony
	BSoissons苏瓦松() feud.Barony
}

type 亚眠AmiensCounty struct {
	feud.BaseCounty
	亚眠Amiens       feud.Barony
	布勒特伊Breteuil   feud.Barony
	科尔比Corbie      feud.Barony
	蒙迪迪耶Montdidier feud.Barony
	内勒Nesle        feud.Barony
	努瓦永Noyon       feud.Barony
	佩罗讷Peronne     feud.Barony
	苏瓦松Soissons    feud.Barony
}

func (c *亚眠AmiensCounty) BAmiens亚眠() feud.Barony {
	return c.亚眠Amiens
}

func (c *亚眠AmiensCounty) BBreteuil布勒特伊() feud.Barony {
	return c.布勒特伊Breteuil
}

func (c *亚眠AmiensCounty) BCorbie科尔比() feud.Barony {
	return c.科尔比Corbie
}

func (c *亚眠AmiensCounty) BMontdidier蒙迪迪耶() feud.Barony {
	return c.蒙迪迪耶Montdidier
}

func (c *亚眠AmiensCounty) BNesle内勒() feud.Barony {
	return c.内勒Nesle
}

func (c *亚眠AmiensCounty) BNoyon努瓦永() feud.Barony {
	return c.努瓦永Noyon
}

func (c *亚眠AmiensCounty) BPeronne佩罗讷() feud.Barony {
	return c.佩罗讷Peronne
}

func (c *亚眠AmiensCounty) BSoissons苏瓦松() feud.Barony {
	return c.苏瓦松Soissons
}

var CAmiens亚眠 AmiensCounty = &亚眠AmiensCounty{}

func init() {
	f := CAmiens亚眠.(*亚眠AmiensCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "95",
		Title:     "amiens",
		TitleName: "亚眠",
		TitleCode: "c_amiens",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚眠Amiens = BAmiens亚眠
	f.亚眠Amiens.SetParent(f)

	f.布勒特伊Breteuil = BBreteuil布勒特伊
	f.布勒特伊Breteuil.SetParent(f)

	f.科尔比Corbie = BCorbie科尔比
	f.科尔比Corbie.SetParent(f)

	f.蒙迪迪耶Montdidier = BMontdidier蒙迪迪耶
	f.蒙迪迪耶Montdidier.SetParent(f)

	f.内勒Nesle = BNesle内勒
	f.内勒Nesle.SetParent(f)

	f.努瓦永Noyon = BNoyon努瓦永
	f.努瓦永Noyon.SetParent(f)

	f.佩罗讷Peronne = BPeronne佩罗讷
	f.佩罗讷Peronne.SetParent(f)

	f.苏瓦松Soissons = BSoissons苏瓦松
	f.苏瓦松Soissons.SetParent(f)

}
