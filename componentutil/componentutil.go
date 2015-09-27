package componentutil

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
	"github.com/neelance/dom/domutil"
)

type Lifecycle interface {
	Render() dom.Spec
}

type Core struct {
	Lifecycle Lifecycle
	Body      dom.Spec
}

func (a *Core) Node() *js.Object {
	return a.Body.Node()
}

func (a *Core) DoRender() {
	oldBody := a.Body
	a.Body = a.Lifecycle.Render()
	a.Body.Reconcile(oldBody)
}

func (a *Core) Update() {
	oldBody := a.Body
	a.DoRender()
	domutil.ReplaceNode(a.Body.Node(), oldBody.Node())
}
