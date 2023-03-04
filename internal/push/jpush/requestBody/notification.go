package requestBody

import (
	"Open_IM/pkg/common/config"
)

type Notification struct {
	Alert   string  `json:"alert,omitempty"`
	Android Android `json:"android,omitempty"`
	IOS     Ios     `json:"ios,omitempty"`
}

type Android struct {
	Alert  string `json:"alert,omitempty"`
	Title  string `json:"title,omitempty"`
	Intent struct {
		URL string `json:"url,omitempty"`
	} `json:"intent,omitempty"`
	Extras Extras `json:"extras"`
}
type Ios struct {
	Alert struct {
		Title string `json:"title,omitempty"`
		Body  string `json:"body,omitempty"`
	} `json:"alert,omitempty"`
	Sound          string `json:"sound,omitempty"`
	Badge          string `json:"badge,omitempty"`
	Extras         Extras `json:"extras"`
	MutableContent bool   `json:"mutable-content"`
}

type Extras struct {
	ClientMsgID string `json:"clientMsgID"`
}

func (n *Notification) SetAlert(title string, alert string) {
	n.Alert = alert
	n.Android.Alert = alert
	n.Android.Title = title
	n.SetAndroidIntent()
	n.IOS.Alert.Body = alert
	n.IOS.Alert.Title = title

	n.IOS.Badge = "+1"
}

func (n *Notification) SetExtras(extras Extras) {
	n.IOS.Extras = extras
	n.Android.Extras = extras
}

func (n *Notification) SetAndroidIntent() {
	n.Android.Intent.URL = config.Config.Push.Jpns.PushIntent
}

func (n *Notification) IOSEnableMutableContent() {
	n.IOS.MutableContent = true
}
