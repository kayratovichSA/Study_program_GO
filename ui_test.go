package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"strings"
	"testing"
)

func TestTabsAndCalculator(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	root := buildUI()
	tabs := root.Objects[0].(*container.AppTabs)
	if len(tabs.Items) != 4 {
		t.Fatal("expected four chapter tabs")
	}
	for _, tab := range tabs.Items {
		tabs.Select(tab)
		if tabs.Selected() != tab {
			t.Fatal("tab selection")
		}
		cards := tab.Content.(*container.Scroll).Content.(*fyne.Container)
		for _, object := range cards.Objects {
			content := object.(*widget.Card).Content.(*fyne.Container)
			if button, ok := content.Objects[2].(*widget.Button); ok {
				test.Tap(button)
				output := content.Objects[3].(*widget.Label).Text
				if output == "" || output == "Результат появится здесь." || strings.HasPrefix(output, "Ошибка:") {
					t.Fatal("task did not produce output")
				}
			}
		}
	}
	scroll := tabs.Items[0].Content.(*container.Scroll)
	card := scroll.Content.(*fyne.Container).Objects[0].(*widget.Card)
	content := card.Content.(*fyne.Container)
	test.Tap(content.Objects[2].(*widget.Button))
	if content.Objects[3].(*widget.Label).Text != "12 + 3 = 15" {
		t.Fatal("calculator button output")
	}
	inputs := content.Objects[1].(*widget.Form)
	inputs.Items[1].Widget.(*widget.Select).SetSelected("÷")
	inputs.Items[2].Widget.(*widget.Entry).SetText("0")
	test.Tap(content.Objects[2].(*widget.Button))
	if !strings.HasPrefix(content.Objects[3].(*widget.Label).Text, "Ошибка:") {
		t.Fatal("division error not displayed")
	}
	counter := tabs.Items[1].Content.(*container.Scroll).Content.(*fyne.Container).Objects[3].(*widget.Card).Content.(*fyne.Container)
	buttons := counter.Objects[1].(*fyne.Container)
	test.Tap(buttons.Objects[0].(*widget.Button))
	test.Tap(buttons.Objects[0].(*widget.Button))
	if counter.Objects[2].(*widget.Label).Text != "Счётчик: 2" {
		t.Fatal("counter")
	}
	test.Tap(buttons.Objects[1].(*widget.Button))
	if counter.Objects[2].(*widget.Label).Text != "Счётчик: 0" {
		t.Fatal("counter reset")
	}
	directory := tabs.Items[3].Content.(*container.Scroll).Content.(*fyne.Container).Objects[3].(*widget.Card).Content.(*fyne.Container)
	controls := directory.Objects[2].(*fyne.Container)
	test.Tap(controls.Objects[0].(*widget.Button))
	test.Tap(controls.Objects[1].(*widget.Button))
	if !strings.Contains(directory.Objects[3].(*widget.Label).Text, "Анна:") {
		t.Fatal("map save/search")
	}
	test.Tap(controls.Objects[2].(*widget.Button))
	if directory.Objects[3].(*widget.Label).Text != "Справочник пуст." {
		t.Fatal("map delete")
	}
}
