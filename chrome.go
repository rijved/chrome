package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	CHROME = "chrome"
)

type Chrome struct {
	o             js.Object
	Tabs          Tabs
	Windows       Windows
	Runtime       Runtime
	Alarms        Alarms
	Bookmarks     Bookmarks
	BrowserAction BrowserAction
	BrowsingData  BrowsingData
	Commands      Commands
	ContextMenus  ContextMenus
}

func NewChrome() Chrome {
	c := Chrome{o: js.Global.Get(CHROME)}
	c.Tabs = Tabs{o: c.o.Get("tabs")}
	c.Windows = Windows{o: c.o.Get("windows")}
	c.Runtime = Runtime{o: c.o.Get("runtime")}
	c.Alarms = Alarms{o: c.o.Get("alarms")}
	c.Bookmarks = Bookmarks{o: c.o.Get("bookmarks")}
	c.BrowserAction = BrowserAction{o: c.o.Get("browserAction")}
	c.BrowsingData = BrowsingData{o: c.o.Get("browsingData")}
	c.Commands = Commands{o: c.o.Get("commands")}
	c.ContextMenus = ContextMenus{o: c.o.Get("contextMenus")}
	return c
}