package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Sarasvata_mandalaCounty interface {
    feud.County
    BAnahilapataka阿那醯罗波吒迦() 	feud.Barony
    BKamboika甘波迦() 	feud.Barony
    BKheralu诃罗卢() 	feud.Barony
    BModhera摩诃罗() 	feud.Barony
    BShertha舍陀() 	feud.Barony
    BSiddhapura悉陀补罗() 	feud.Barony
    BVadnagar跋那揭罗() 	feud.Barony
}

type 萨罗萨伐多曼荼罗Sarasvata_mandalaCounty struct {
	feud.BaseCounty
	阿那醯罗波吒迦Anahilapataka 	feud.Barony
	甘波迦Kamboika 	feud.Barony
	诃罗卢Kheralu 	feud.Barony
	摩诃罗Modhera 	feud.Barony
	舍陀Shertha 	feud.Barony
	悉陀补罗Siddhapura 	feud.Barony
	跋那揭罗Vadnagar 	feud.Barony
}

func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BAnahilapataka阿那醯罗波吒迦() feud.Barony {
	return c.阿那醯罗波吒迦Anahilapataka
}
    
func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BKamboika甘波迦() feud.Barony {
	return c.甘波迦Kamboika
}
    
func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BKheralu诃罗卢() feud.Barony {
	return c.诃罗卢Kheralu
}
    
func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BModhera摩诃罗() feud.Barony {
	return c.摩诃罗Modhera
}
    
func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BShertha舍陀() feud.Barony {
	return c.舍陀Shertha
}
    
func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BSiddhapura悉陀补罗() feud.Barony {
	return c.悉陀补罗Siddhapura
}
    
func (c *萨罗萨伐多曼荼罗Sarasvata_mandalaCounty) BVadnagar跋那揭罗() feud.Barony {
	return c.跋那揭罗Vadnagar
}
    
var CSarasvata_mandala萨罗萨伐多曼荼罗 Sarasvata_mandalaCounty = &萨罗萨伐多曼荼罗Sarasvata_mandalaCounty{}

func init() {
	f := CSarasvata_mandala萨罗萨伐多曼荼罗.(*萨罗萨伐多曼荼罗Sarasvata_mandalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1292",
		Title:     "sarasvata_mandala",
		TitleName: "萨罗萨伐多曼荼罗",
		TitleCode: "c_sarasvata_mandala",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿那醯罗波吒迦Anahilapataka = BAnahilapataka阿那醯罗波吒迦
	f.阿那醯罗波吒迦Anahilapataka.SetParent(f)
	
	f.甘波迦Kamboika = BKamboika甘波迦
	f.甘波迦Kamboika.SetParent(f)
	
	f.诃罗卢Kheralu = BKheralu诃罗卢
	f.诃罗卢Kheralu.SetParent(f)
	
	f.摩诃罗Modhera = BModhera摩诃罗
	f.摩诃罗Modhera.SetParent(f)
	
	f.舍陀Shertha = BShertha舍陀
	f.舍陀Shertha.SetParent(f)
	
	f.悉陀补罗Siddhapura = BSiddhapura悉陀补罗
	f.悉陀补罗Siddhapura.SetParent(f)
	
	f.跋那揭罗Vadnagar = BVadnagar跋那揭罗
	f.跋那揭罗Vadnagar.SetParent(f)
	
}
