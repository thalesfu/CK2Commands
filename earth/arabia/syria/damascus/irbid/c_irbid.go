package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IrbidCounty interface {
    feud.County
    BAljun阿杰隆() 	feud.Barony
    BAydoun埃东尼() 	feud.Barony
    BGadera加大拉() 	feud.Barony
    BHabisjaldak哈比斯贾勒达克() 	feud.Barony
    BIrbid伊尔比德() 	feud.Barony
    BPella佩拉() 	feud.Barony
    BUmmqais乌姆盖斯() 	feud.Barony
    BYarmouk耶尔穆克() 	feud.Barony
}

type 伊尔比德IrbidCounty struct {
	feud.BaseCounty
	阿杰隆Aljun 	feud.Barony
	埃东尼Aydoun 	feud.Barony
	加大拉Gadera 	feud.Barony
	哈比斯贾勒达克Habisjaldak 	feud.Barony
	伊尔比德Irbid 	feud.Barony
	佩拉Pella 	feud.Barony
	乌姆盖斯Ummqais 	feud.Barony
	耶尔穆克Yarmouk 	feud.Barony
}

func (c *伊尔比德IrbidCounty) BAljun阿杰隆() feud.Barony {
	return c.阿杰隆Aljun
}
    
func (c *伊尔比德IrbidCounty) BAydoun埃东尼() feud.Barony {
	return c.埃东尼Aydoun
}
    
func (c *伊尔比德IrbidCounty) BGadera加大拉() feud.Barony {
	return c.加大拉Gadera
}
    
func (c *伊尔比德IrbidCounty) BHabisjaldak哈比斯贾勒达克() feud.Barony {
	return c.哈比斯贾勒达克Habisjaldak
}
    
func (c *伊尔比德IrbidCounty) BIrbid伊尔比德() feud.Barony {
	return c.伊尔比德Irbid
}
    
func (c *伊尔比德IrbidCounty) BPella佩拉() feud.Barony {
	return c.佩拉Pella
}
    
func (c *伊尔比德IrbidCounty) BUmmqais乌姆盖斯() feud.Barony {
	return c.乌姆盖斯Ummqais
}
    
func (c *伊尔比德IrbidCounty) BYarmouk耶尔穆克() feud.Barony {
	return c.耶尔穆克Yarmouk
}
    
var CIrbid伊尔比德 IrbidCounty = &伊尔比德IrbidCounty{}

func init() {
	f := CIrbid伊尔比德.(*伊尔比德IrbidCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "725",
		Title:     "irbid",
		TitleName: "伊尔比德",
		TitleCode: "c_irbid",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿杰隆Aljun = BAljun阿杰隆
	f.阿杰隆Aljun.SetParent(f)
	
	f.埃东尼Aydoun = BAydoun埃东尼
	f.埃东尼Aydoun.SetParent(f)
	
	f.加大拉Gadera = BGadera加大拉
	f.加大拉Gadera.SetParent(f)
	
	f.哈比斯贾勒达克Habisjaldak = BHabisjaldak哈比斯贾勒达克
	f.哈比斯贾勒达克Habisjaldak.SetParent(f)
	
	f.伊尔比德Irbid = BIrbid伊尔比德
	f.伊尔比德Irbid.SetParent(f)
	
	f.佩拉Pella = BPella佩拉
	f.佩拉Pella.SetParent(f)
	
	f.乌姆盖斯Ummqais = BUmmqais乌姆盖斯
	f.乌姆盖斯Ummqais.SetParent(f)
	
	f.耶尔穆克Yarmouk = BYarmouk耶尔穆克
	f.耶尔穆克Yarmouk.SetParent(f)
	
}
