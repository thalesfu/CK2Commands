package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LusignanCounty interface {
    feud.County
    BCharroux沙鲁() 	feud.Barony
    BCivray锡夫赖() 	feud.Barony
    BConfolens孔福朗() 	feud.Barony
    BLusignan吕西尼昂() 	feud.Barony
    BMaillezais马耶赛() 	feud.Barony
    BMelle梅勒() 	feud.Barony
    BNiort尼奥尔() 	feud.Barony
    BStmaixent圣迈克桑() 	feud.Barony
}

type 吕西尼昂LusignanCounty struct {
	feud.BaseCounty
	沙鲁Charroux 	feud.Barony
	锡夫赖Civray 	feud.Barony
	孔福朗Confolens 	feud.Barony
	吕西尼昂Lusignan 	feud.Barony
	马耶赛Maillezais 	feud.Barony
	梅勒Melle 	feud.Barony
	尼奥尔Niort 	feud.Barony
	圣迈克桑Stmaixent 	feud.Barony
}

func (c *吕西尼昂LusignanCounty) BCharroux沙鲁() feud.Barony {
	return c.沙鲁Charroux
}
    
func (c *吕西尼昂LusignanCounty) BCivray锡夫赖() feud.Barony {
	return c.锡夫赖Civray
}
    
func (c *吕西尼昂LusignanCounty) BConfolens孔福朗() feud.Barony {
	return c.孔福朗Confolens
}
    
func (c *吕西尼昂LusignanCounty) BLusignan吕西尼昂() feud.Barony {
	return c.吕西尼昂Lusignan
}
    
func (c *吕西尼昂LusignanCounty) BMaillezais马耶赛() feud.Barony {
	return c.马耶赛Maillezais
}
    
func (c *吕西尼昂LusignanCounty) BMelle梅勒() feud.Barony {
	return c.梅勒Melle
}
    
func (c *吕西尼昂LusignanCounty) BNiort尼奥尔() feud.Barony {
	return c.尼奥尔Niort
}
    
func (c *吕西尼昂LusignanCounty) BStmaixent圣迈克桑() feud.Barony {
	return c.圣迈克桑Stmaixent
}
    
var CLusignan吕西尼昂 LusignanCounty = &吕西尼昂LusignanCounty{}

func init() {
	f := CLusignan吕西尼昂.(*吕西尼昂LusignanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "144",
		Title:     "lusignan",
		TitleName: "吕西尼昂",
		TitleCode: "c_lusignan",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙鲁Charroux = BCharroux沙鲁
	f.沙鲁Charroux.SetParent(f)
	
	f.锡夫赖Civray = BCivray锡夫赖
	f.锡夫赖Civray.SetParent(f)
	
	f.孔福朗Confolens = BConfolens孔福朗
	f.孔福朗Confolens.SetParent(f)
	
	f.吕西尼昂Lusignan = BLusignan吕西尼昂
	f.吕西尼昂Lusignan.SetParent(f)
	
	f.马耶赛Maillezais = BMaillezais马耶赛
	f.马耶赛Maillezais.SetParent(f)
	
	f.梅勒Melle = BMelle梅勒
	f.梅勒Melle.SetParent(f)
	
	f.尼奥尔Niort = BNiort尼奥尔
	f.尼奥尔Niort.SetParent(f)
	
	f.圣迈克桑Stmaixent = BStmaixent圣迈克桑
	f.圣迈克桑Stmaixent.SetParent(f)
	
}
