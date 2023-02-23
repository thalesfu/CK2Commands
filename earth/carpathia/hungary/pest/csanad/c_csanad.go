package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CsanadCounty interface {
	feud.County
	BBattonya鲍托尼奥() feud.Barony
	BCsanad乔纳德() feud.Barony
	BCsongrad琼格拉德() feud.Barony
	BHodmezovasarhely霍德迈泽瓦沙尔海伊() feud.Barony
	BMako毛科() feud.Barony
	BMindszent明曾特() feud.Barony
	BSzeged塞盖德() feud.Barony
	BSzentes塞内泰什() feud.Barony
}

type 乔纳德CsanadCounty struct {
	feud.BaseCounty
	鲍托尼奥Battonya              feud.Barony
	乔纳德Csanad                 feud.Barony
	琼格拉德Csongrad              feud.Barony
	霍德迈泽瓦沙尔海伊Hodmezovasarhely feud.Barony
	毛科Mako                    feud.Barony
	明曾特Mindszent              feud.Barony
	塞盖德Szeged                 feud.Barony
	塞内泰什Szentes               feud.Barony
}

func (c *乔纳德CsanadCounty) BBattonya鲍托尼奥() feud.Barony {
	return c.鲍托尼奥Battonya
}

func (c *乔纳德CsanadCounty) BCsanad乔纳德() feud.Barony {
	return c.乔纳德Csanad
}

func (c *乔纳德CsanadCounty) BCsongrad琼格拉德() feud.Barony {
	return c.琼格拉德Csongrad
}

func (c *乔纳德CsanadCounty) BHodmezovasarhely霍德迈泽瓦沙尔海伊() feud.Barony {
	return c.霍德迈泽瓦沙尔海伊Hodmezovasarhely
}

func (c *乔纳德CsanadCounty) BMako毛科() feud.Barony {
	return c.毛科Mako
}

func (c *乔纳德CsanadCounty) BMindszent明曾特() feud.Barony {
	return c.明曾特Mindszent
}

func (c *乔纳德CsanadCounty) BSzeged塞盖德() feud.Barony {
	return c.塞盖德Szeged
}

func (c *乔纳德CsanadCounty) BSzentes塞内泰什() feud.Barony {
	return c.塞内泰什Szentes
}

var CCsanad乔纳德 CsanadCounty = &乔纳德CsanadCounty{}

func init() {
	f := CCsanad乔纳德.(*乔纳德CsanadCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "521",
		Title:     "csanad",
		TitleName: "乔纳德",
		TitleCode: "c_csanad",
		Baronies:  map[string]feud.Barony{},
	}

	f.鲍托尼奥Battonya = BBattonya鲍托尼奥
	f.鲍托尼奥Battonya.SetParent(f)

	f.乔纳德Csanad = BCsanad乔纳德
	f.乔纳德Csanad.SetParent(f)

	f.琼格拉德Csongrad = BCsongrad琼格拉德
	f.琼格拉德Csongrad.SetParent(f)

	f.霍德迈泽瓦沙尔海伊Hodmezovasarhely = BHodmezovasarhely霍德迈泽瓦沙尔海伊
	f.霍德迈泽瓦沙尔海伊Hodmezovasarhely.SetParent(f)

	f.毛科Mako = BMako毛科
	f.毛科Mako.SetParent(f)

	f.明曾特Mindszent = BMindszent明曾特
	f.明曾特Mindszent.SetParent(f)

	f.塞盖德Szeged = BSzeged塞盖德
	f.塞盖德Szeged.SetParent(f)

	f.塞内泰什Szentes = BSzentes塞内泰什
	f.塞内泰什Szentes.SetParent(f)

}
