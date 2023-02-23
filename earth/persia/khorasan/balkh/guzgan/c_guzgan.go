package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GuzganCounty interface {
	feud.County
	BAndkhud安狄枯() feud.Barony
	BFaryab法里亚布() feud.Barony
	BJahudan雅胡丹() feud.Barony
	BNaryan那里杨() feud.Barony
	BQasrahnaf阿希纳弗堡() feud.Barony
	BTalaqan呾剌健() feud.Barony
	BUsbuman乌斯卜曼() feud.Barony
}

type 胡实健GuzganCounty struct {
	feud.BaseCounty
	安狄枯Andkhud     feud.Barony
	法里亚布Faryab     feud.Barony
	雅胡丹Jahudan     feud.Barony
	那里杨Naryan      feud.Barony
	阿希纳弗堡Qasrahnaf feud.Barony
	呾剌健Talaqan     feud.Barony
	乌斯卜曼Usbuman    feud.Barony
}

func (c *胡实健GuzganCounty) BAndkhud安狄枯() feud.Barony {
	return c.安狄枯Andkhud
}

func (c *胡实健GuzganCounty) BFaryab法里亚布() feud.Barony {
	return c.法里亚布Faryab
}

func (c *胡实健GuzganCounty) BJahudan雅胡丹() feud.Barony {
	return c.雅胡丹Jahudan
}

func (c *胡实健GuzganCounty) BNaryan那里杨() feud.Barony {
	return c.那里杨Naryan
}

func (c *胡实健GuzganCounty) BQasrahnaf阿希纳弗堡() feud.Barony {
	return c.阿希纳弗堡Qasrahnaf
}

func (c *胡实健GuzganCounty) BTalaqan呾剌健() feud.Barony {
	return c.呾剌健Talaqan
}

func (c *胡实健GuzganCounty) BUsbuman乌斯卜曼() feud.Barony {
	return c.乌斯卜曼Usbuman
}

var CGuzgan胡实健 GuzganCounty = &胡实健GuzganCounty{}

func init() {
	f := CGuzgan胡实健.(*胡实健GuzganCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1540",
		Title:     "guzgan",
		TitleName: "胡实健",
		TitleCode: "c_guzgan",
		Baronies:  map[string]feud.Barony{},
	}

	f.安狄枯Andkhud = BAndkhud安狄枯
	f.安狄枯Andkhud.SetParent(f)

	f.法里亚布Faryab = BFaryab法里亚布
	f.法里亚布Faryab.SetParent(f)

	f.雅胡丹Jahudan = BJahudan雅胡丹
	f.雅胡丹Jahudan.SetParent(f)

	f.那里杨Naryan = BNaryan那里杨
	f.那里杨Naryan.SetParent(f)

	f.阿希纳弗堡Qasrahnaf = BQasrahnaf阿希纳弗堡
	f.阿希纳弗堡Qasrahnaf.SetParent(f)

	f.呾剌健Talaqan = BTalaqan呾剌健
	f.呾剌健Talaqan.SetParent(f)

	f.乌斯卜曼Usbuman = BUsbuman乌斯卜曼
	f.乌斯卜曼Usbuman.SetParent(f)

}
