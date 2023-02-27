package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarakhsCounty interface {
    feud.County
    BDandanqan丹丹坎() 	feud.Barony
    BMaihana迈哈娜() 	feud.Barony
    BMazduran马兹杜兰() 	feud.Barony
    BPaxtachi帕赫塔基() 	feud.Barony
    BSarakhs撒剌哈夕() 	feud.Barony
    BShiraz_merv设拉子() 	feud.Barony
    BYuzlar尤扎拉() 	feud.Barony
}

type 撒剌哈夕SarakhsCounty struct {
	feud.BaseCounty
	丹丹坎Dandanqan 	feud.Barony
	迈哈娜Maihana 	feud.Barony
	马兹杜兰Mazduran 	feud.Barony
	帕赫塔基Paxtachi 	feud.Barony
	撒剌哈夕Sarakhs 	feud.Barony
	设拉子Shiraz_merv 	feud.Barony
	尤扎拉Yuzlar 	feud.Barony
}

func (c *撒剌哈夕SarakhsCounty) BDandanqan丹丹坎() feud.Barony {
	return c.丹丹坎Dandanqan
}
    
func (c *撒剌哈夕SarakhsCounty) BMaihana迈哈娜() feud.Barony {
	return c.迈哈娜Maihana
}
    
func (c *撒剌哈夕SarakhsCounty) BMazduran马兹杜兰() feud.Barony {
	return c.马兹杜兰Mazduran
}
    
func (c *撒剌哈夕SarakhsCounty) BPaxtachi帕赫塔基() feud.Barony {
	return c.帕赫塔基Paxtachi
}
    
func (c *撒剌哈夕SarakhsCounty) BSarakhs撒剌哈夕() feud.Barony {
	return c.撒剌哈夕Sarakhs
}
    
func (c *撒剌哈夕SarakhsCounty) BShiraz_merv设拉子() feud.Barony {
	return c.设拉子Shiraz_merv
}
    
func (c *撒剌哈夕SarakhsCounty) BYuzlar尤扎拉() feud.Barony {
	return c.尤扎拉Yuzlar
}
    
var CSarakhs撒剌哈夕 SarakhsCounty = &撒剌哈夕SarakhsCounty{}

func init() {
	f := CSarakhs撒剌哈夕.(*撒剌哈夕SarakhsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1542",
		Title:     "sarakhs",
		TitleName: "撒剌哈夕",
		TitleCode: "c_sarakhs",
		Baronies:  map[string]feud.Barony{},
	}

	f.丹丹坎Dandanqan = BDandanqan丹丹坎
	f.丹丹坎Dandanqan.SetParent(f)
	
	f.迈哈娜Maihana = BMaihana迈哈娜
	f.迈哈娜Maihana.SetParent(f)
	
	f.马兹杜兰Mazduran = BMazduran马兹杜兰
	f.马兹杜兰Mazduran.SetParent(f)
	
	f.帕赫塔基Paxtachi = BPaxtachi帕赫塔基
	f.帕赫塔基Paxtachi.SetParent(f)
	
	f.撒剌哈夕Sarakhs = BSarakhs撒剌哈夕
	f.撒剌哈夕Sarakhs.SetParent(f)
	
	f.设拉子Shiraz_merv = BShiraz_merv设拉子
	f.设拉子Shiraz_merv.SetParent(f)
	
	f.尤扎拉Yuzlar = BYuzlar尤扎拉
	f.尤扎拉Yuzlar.SetParent(f)
	
}
