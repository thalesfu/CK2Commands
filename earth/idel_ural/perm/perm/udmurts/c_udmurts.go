package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UdmurtsCounty interface {
    feud.County
    BIz伊兹() 	feud.Barony
    BKarlutka卡尔卢特卡() 	feud.Barony
    BPostol波斯托尔() 	feud.Barony
    BSarkan萨尔坎() 	feud.Barony
    BUdmurts乌德穆尔特() 	feud.Barony
    BUva乌瓦() 	feud.Barony
    BVotkinsk沃特金斯克() 	feud.Barony
}

type 乌德穆尔特UdmurtsCounty struct {
	feud.BaseCounty
	伊兹Iz 	feud.Barony
	卡尔卢特卡Karlutka 	feud.Barony
	波斯托尔Postol 	feud.Barony
	萨尔坎Sarkan 	feud.Barony
	乌德穆尔特Udmurts 	feud.Barony
	乌瓦Uva 	feud.Barony
	沃特金斯克Votkinsk 	feud.Barony
}

func (c *乌德穆尔特UdmurtsCounty) BIz伊兹() feud.Barony {
	return c.伊兹Iz
}
    
func (c *乌德穆尔特UdmurtsCounty) BKarlutka卡尔卢特卡() feud.Barony {
	return c.卡尔卢特卡Karlutka
}
    
func (c *乌德穆尔特UdmurtsCounty) BPostol波斯托尔() feud.Barony {
	return c.波斯托尔Postol
}
    
func (c *乌德穆尔特UdmurtsCounty) BSarkan萨尔坎() feud.Barony {
	return c.萨尔坎Sarkan
}
    
func (c *乌德穆尔特UdmurtsCounty) BUdmurts乌德穆尔特() feud.Barony {
	return c.乌德穆尔特Udmurts
}
    
func (c *乌德穆尔特UdmurtsCounty) BUva乌瓦() feud.Barony {
	return c.乌瓦Uva
}
    
func (c *乌德穆尔特UdmurtsCounty) BVotkinsk沃特金斯克() feud.Barony {
	return c.沃特金斯克Votkinsk
}
    
var CUdmurts乌德穆尔特 UdmurtsCounty = &乌德穆尔特UdmurtsCounty{}

func init() {
	f := CUdmurts乌德穆尔特.(*乌德穆尔特UdmurtsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1719",
		Title:     "udmurts",
		TitleName: "乌德穆尔特",
		TitleCode: "c_udmurts",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊兹Iz = BIz伊兹
	f.伊兹Iz.SetParent(f)
	
	f.卡尔卢特卡Karlutka = BKarlutka卡尔卢特卡
	f.卡尔卢特卡Karlutka.SetParent(f)
	
	f.波斯托尔Postol = BPostol波斯托尔
	f.波斯托尔Postol.SetParent(f)
	
	f.萨尔坎Sarkan = BSarkan萨尔坎
	f.萨尔坎Sarkan.SetParent(f)
	
	f.乌德穆尔特Udmurts = BUdmurts乌德穆尔特
	f.乌德穆尔特Udmurts.SetParent(f)
	
	f.乌瓦Uva = BUva乌瓦
	f.乌瓦Uva.SetParent(f)
	
	f.沃特金斯克Votkinsk = BVotkinsk沃特金斯克
	f.沃特金斯克Votkinsk.SetParent(f)
	
}
