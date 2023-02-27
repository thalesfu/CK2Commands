package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DelhiCounty interface {
    feud.County
    BDhilika提梨迦() 	feud.Barony
    BIndraprastha因陀罗钵罗萨他() 	feud.Barony
    BJahanpanah贾汉帕那赫() 	feud.Barony
    BLalkot剌罗拘吒() 	feud.Barony
    BMehrauli咩诃罗郁梨() 	feud.Barony
    BNarela那梨罗() 	feud.Barony
    BSiri斯利() 	feud.Barony
}

type 德里DelhiCounty struct {
	feud.BaseCounty
	提梨迦Dhilika 	feud.Barony
	因陀罗钵罗萨他Indraprastha 	feud.Barony
	贾汉帕那赫Jahanpanah 	feud.Barony
	剌罗拘吒Lalkot 	feud.Barony
	咩诃罗郁梨Mehrauli 	feud.Barony
	那梨罗Narela 	feud.Barony
	斯利Siri 	feud.Barony
}

func (c *德里DelhiCounty) BDhilika提梨迦() feud.Barony {
	return c.提梨迦Dhilika
}
    
func (c *德里DelhiCounty) BIndraprastha因陀罗钵罗萨他() feud.Barony {
	return c.因陀罗钵罗萨他Indraprastha
}
    
func (c *德里DelhiCounty) BJahanpanah贾汉帕那赫() feud.Barony {
	return c.贾汉帕那赫Jahanpanah
}
    
func (c *德里DelhiCounty) BLalkot剌罗拘吒() feud.Barony {
	return c.剌罗拘吒Lalkot
}
    
func (c *德里DelhiCounty) BMehrauli咩诃罗郁梨() feud.Barony {
	return c.咩诃罗郁梨Mehrauli
}
    
func (c *德里DelhiCounty) BNarela那梨罗() feud.Barony {
	return c.那梨罗Narela
}
    
func (c *德里DelhiCounty) BSiri斯利() feud.Barony {
	return c.斯利Siri
}
    
var CDelhi德里 DelhiCounty = &德里DelhiCounty{}

func init() {
	f := CDelhi德里.(*德里DelhiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1365",
		Title:     "delhi",
		TitleName: "德里",
		TitleCode: "c_delhi",
		Baronies:  map[string]feud.Barony{},
	}

	f.提梨迦Dhilika = BDhilika提梨迦
	f.提梨迦Dhilika.SetParent(f)
	
	f.因陀罗钵罗萨他Indraprastha = BIndraprastha因陀罗钵罗萨他
	f.因陀罗钵罗萨他Indraprastha.SetParent(f)
	
	f.贾汉帕那赫Jahanpanah = BJahanpanah贾汉帕那赫
	f.贾汉帕那赫Jahanpanah.SetParent(f)
	
	f.剌罗拘吒Lalkot = BLalkot剌罗拘吒
	f.剌罗拘吒Lalkot.SetParent(f)
	
	f.咩诃罗郁梨Mehrauli = BMehrauli咩诃罗郁梨
	f.咩诃罗郁梨Mehrauli.SetParent(f)
	
	f.那梨罗Narela = BNarela那梨罗
	f.那梨罗Narela.SetParent(f)
	
	f.斯利Siri = BSiri斯利
	f.斯利Siri.SetParent(f)
	
}
