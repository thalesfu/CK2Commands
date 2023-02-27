package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TavdaCounty interface {
    feud.County
    BBaykalovo拜卡洛沃() 	feud.Barony
    BIgovskoy伊戈夫斯科伊() 	feud.Barony
    BTalitsa塔利察() 	feud.Barony
    BTavda塔夫达() 	feud.Barony
    BTugulym_tavda图古雷姆() 	feud.Barony
    BYar亚尔() 	feud.Barony
    BYushala尤沙拉() 	feud.Barony
}

type 塔夫达TavdaCounty struct {
	feud.BaseCounty
	拜卡洛沃Baykalovo 	feud.Barony
	伊戈夫斯科伊Igovskoy 	feud.Barony
	塔利察Talitsa 	feud.Barony
	塔夫达Tavda 	feud.Barony
	图古雷姆Tugulym_tavda 	feud.Barony
	亚尔Yar 	feud.Barony
	尤沙拉Yushala 	feud.Barony
}

func (c *塔夫达TavdaCounty) BBaykalovo拜卡洛沃() feud.Barony {
	return c.拜卡洛沃Baykalovo
}
    
func (c *塔夫达TavdaCounty) BIgovskoy伊戈夫斯科伊() feud.Barony {
	return c.伊戈夫斯科伊Igovskoy
}
    
func (c *塔夫达TavdaCounty) BTalitsa塔利察() feud.Barony {
	return c.塔利察Talitsa
}
    
func (c *塔夫达TavdaCounty) BTavda塔夫达() feud.Barony {
	return c.塔夫达Tavda
}
    
func (c *塔夫达TavdaCounty) BTugulym_tavda图古雷姆() feud.Barony {
	return c.图古雷姆Tugulym_tavda
}
    
func (c *塔夫达TavdaCounty) BYar亚尔() feud.Barony {
	return c.亚尔Yar
}
    
func (c *塔夫达TavdaCounty) BYushala尤沙拉() feud.Barony {
	return c.尤沙拉Yushala
}
    
var CTavda塔夫达 TavdaCounty = &塔夫达TavdaCounty{}

func init() {
	f := CTavda塔夫达.(*塔夫达TavdaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1858",
		Title:     "tavda",
		TitleName: "塔夫达",
		TitleCode: "c_tavda",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜卡洛沃Baykalovo = BBaykalovo拜卡洛沃
	f.拜卡洛沃Baykalovo.SetParent(f)
	
	f.伊戈夫斯科伊Igovskoy = BIgovskoy伊戈夫斯科伊
	f.伊戈夫斯科伊Igovskoy.SetParent(f)
	
	f.塔利察Talitsa = BTalitsa塔利察
	f.塔利察Talitsa.SetParent(f)
	
	f.塔夫达Tavda = BTavda塔夫达
	f.塔夫达Tavda.SetParent(f)
	
	f.图古雷姆Tugulym_tavda = BTugulym_tavda图古雷姆
	f.图古雷姆Tugulym_tavda.SetParent(f)
	
	f.亚尔Yar = BYar亚尔
	f.亚尔Yar.SetParent(f)
	
	f.尤沙拉Yushala = BYushala尤沙拉
	f.尤沙拉Yushala.SetParent(f)
	
}
