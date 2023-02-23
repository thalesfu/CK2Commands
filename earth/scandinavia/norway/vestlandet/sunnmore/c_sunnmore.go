package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SunnmoreCounty interface {
	feud.County
	BAheim奥海姆() feud.Barony
	BBorgund博尔贡() feud.Barony
	BBorgundkaupangen博尔贡凯于庞恩() feud.Barony
	BGiske伊斯克() feud.Barony
	BHaram哈拉姆() feud.Barony
	BSteinvag斯泰因沃格() feud.Barony
	BValldal瓦尔达尔() feud.Barony
}

type 孙默勒SunnmoreCounty struct {
	feud.BaseCounty
	奥海姆Aheim                feud.Barony
	博尔贡Borgund              feud.Barony
	博尔贡凯于庞恩Borgundkaupangen feud.Barony
	伊斯克Giske                feud.Barony
	哈拉姆Haram                feud.Barony
	斯泰因沃格Steinvag           feud.Barony
	瓦尔达尔Valldal             feud.Barony
}

func (c *孙默勒SunnmoreCounty) BAheim奥海姆() feud.Barony {
	return c.奥海姆Aheim
}

func (c *孙默勒SunnmoreCounty) BBorgund博尔贡() feud.Barony {
	return c.博尔贡Borgund
}

func (c *孙默勒SunnmoreCounty) BBorgundkaupangen博尔贡凯于庞恩() feud.Barony {
	return c.博尔贡凯于庞恩Borgundkaupangen
}

func (c *孙默勒SunnmoreCounty) BGiske伊斯克() feud.Barony {
	return c.伊斯克Giske
}

func (c *孙默勒SunnmoreCounty) BHaram哈拉姆() feud.Barony {
	return c.哈拉姆Haram
}

func (c *孙默勒SunnmoreCounty) BSteinvag斯泰因沃格() feud.Barony {
	return c.斯泰因沃格Steinvag
}

func (c *孙默勒SunnmoreCounty) BValldal瓦尔达尔() feud.Barony {
	return c.瓦尔达尔Valldal
}

var CSunnmore孙默勒 SunnmoreCounty = &孙默勒SunnmoreCounty{}

func init() {
	f := CSunnmore孙默勒.(*孙默勒SunnmoreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1623",
		Title:     "sunnmore",
		TitleName: "孙默勒",
		TitleCode: "c_sunnmore",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥海姆Aheim = BAheim奥海姆
	f.奥海姆Aheim.SetParent(f)

	f.博尔贡Borgund = BBorgund博尔贡
	f.博尔贡Borgund.SetParent(f)

	f.博尔贡凯于庞恩Borgundkaupangen = BBorgundkaupangen博尔贡凯于庞恩
	f.博尔贡凯于庞恩Borgundkaupangen.SetParent(f)

	f.伊斯克Giske = BGiske伊斯克
	f.伊斯克Giske.SetParent(f)

	f.哈拉姆Haram = BHaram哈拉姆
	f.哈拉姆Haram.SetParent(f)

	f.斯泰因沃格Steinvag = BSteinvag斯泰因沃格
	f.斯泰因沃格Steinvag.SetParent(f)

	f.瓦尔达尔Valldal = BValldal瓦尔达尔
	f.瓦尔达尔Valldal.SetParent(f)

}
