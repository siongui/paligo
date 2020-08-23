package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
)

var storageKeyName = "PaliSetting"

func SavePaliSetting(setting lib.PaliSetting) {
	str, _ := lib.PaliSettingToJsonString(setting)
	LocalStorage.Set(storageKeyName, str)
}

func LoadPaliSetting() lib.PaliSetting {
	setting, _ := lib.JsonStringToPaliSetting(LocalStorage.GetItem(storageKeyName))
	return setting
}

func SetupPaliSetting() {
	isPreview := Document.GetElementById("isShowWordPreview")
	p2en := Document.GetElementById("p2en")
	p2ja := Document.GetElementById("p2ja")
	p2zh := Document.GetElementById("p2zh")
	p2vi := Document.GetElementById("p2vi")
	p2my := Document.GetElementById("p2my")
	dicLangOrder := Document.GetElementById("dicLangOrder")

	setting := lib.GetDefaultPaliSetting()
	// check if there is saved setting in user browser
	if LocalStorage.IsKeyExist(storageKeyName) {
		// use saved setting
		setting, _ = lib.JsonStringToPaliSetting(LocalStorage.GetItem(storageKeyName))
	} else {
		// no setting saved, use default setting
		SavePaliSetting(setting)
	}

	// restore setting
	isPreview.Set("checked", setting.IsShowWordPreview)
	p2en.Set("checked", setting.P2en)
	p2ja.Set("checked", setting.P2ja)
	p2zh.Set("checked", setting.P2zh)
	p2vi.Set("checked", setting.P2vi)
	p2my.Set("checked", setting.P2my)
	dicLangOrder.Set("value", setting.DicLangOrder)

	// set up event handler for setting change
	isPreview.AddEventListener("click", func(e Event) {
		setting.IsShowWordPreview = isPreview.Get("checked").Bool()
		SavePaliSetting(setting)
	})
	// https://stackoverflow.com/questions/4471401/getting-value-of-html-checkbox-from-onclick-onchange-events
	p2en.AddEventListener("click", func(e Event) {
		setting.P2en = p2en.Get("checked").Bool()
		SavePaliSetting(setting)
	})
	p2ja.AddEventListener("click", func(e Event) {
		setting.P2ja = p2ja.Get("checked").Bool()
		SavePaliSetting(setting)
	})
	p2zh.AddEventListener("click", func(e Event) {
		setting.P2zh = p2zh.Get("checked").Bool()
		SavePaliSetting(setting)
	})
	p2vi.AddEventListener("click", func(e Event) {
		setting.P2vi = p2vi.Get("checked").Bool()
		SavePaliSetting(setting)
	})
	p2my.AddEventListener("click", func(e Event) {
		setting.P2my = p2my.Get("checked").Bool()
		SavePaliSetting(setting)
	})
	dicLangOrder.AddEventListener("change", func(e Event) {
		setting.DicLangOrder = dicLangOrder.Get("options").Call("item",
			dicLangOrder.Get("selectedIndex").Int()).Get("value").String()
		SavePaliSetting(setting)
	})
}
