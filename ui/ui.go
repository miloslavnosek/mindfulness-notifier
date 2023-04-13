package ui

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"notifier/database"
)

func CreateViewNotificationList(notifications []string) *widget.List {
	list := widget.NewList(
		func() int {
			return len(notifications)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(index widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(notifications[index])
		},
	)

	return list
}

func CreateAddNotificationVBox(db *sql.DB) *fyne.Container {
	notificationAddEntry := widget.NewEntry()
	notificationAddEntry.SetPlaceHolder("Add a notification")

	notificationAddSubmitButton := widget.NewButton(
		"Add", func() {
			database.InsertNotification(db, notificationAddEntry.Text)
		},
	)

	return container.NewVBox(
		notificationAddEntry,
		notificationAddSubmitButton,
	)
}

func CreateSendNotificationToggle(onToggle func()) *widget.Button {
	return widget.NewButton("Toggle sending of notifications", onToggle)
}