package controller

import (
	"github.com/c4dt/dela-apps/gapi"
	guictrl "github.com/c4dt/dela-apps/gapi/controller/gui/controllers"
	"github.com/c4dt/dela/cli/node"
	"github.com/c4dt/dela/core/ordering/cosipbft"
	"github.com/c4dt/dela/mino/minogrpc"
	"github.com/c4dt/dela/mino/proxy"
	"golang.org/x/xerrors"
)

// registerAction is an action that registers the handlers to the dela proxy
//
// - implements node.ActionTemplate
type registerAction struct{}

func (a registerAction) Execute(ctx node.Context) error {
	var cosi *cosipbft.Service
	err := ctx.Injector.Resolve(&cosi)
	if err != nil {
		return xerrors.Errorf("failed to resolve cosi: %v", err)
	}

	var grpc *minogrpc.Minogrpc
	err = ctx.Injector.Resolve(&grpc)
	if err != nil {
		return xerrors.Errorf("failed to resolve mino grpc: %v", err)
	}

	api := gapi.NewGAPI(cosi, grpc)

	ctx.Injector.Inject(&api)

	var proxy proxy.Proxy
	err = ctx.Injector.Resolve(&proxy)
	if err != nil {
		return xerrors.Errorf("failed to resolve proxy: %v", err)
	}

	ctrl := guictrl.NewCtrl(&api)

	proxy.RegisterHandler("/transactions", ctrl.Transaction())
	proxy.RegisterHandler("/store", ctrl.Store())
	proxy.RegisterHandler("/traffic/sent", ctrl.Sent())

	return nil
}
