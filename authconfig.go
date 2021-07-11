package authconfig

import (
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
)

type ABOption func(*authboss.Authboss)

func NewAuth(options ...ABOption) (*authboss.Authboss, error) {
	ab := authboss.New()

	for _, option := range options {
		option(ab)
	}

	return ab, nil
}

func Server(srv authboss.ServerStorer) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Storage.Server = srv
	}
}

func CookieStorer(store abclientstate.CookieStorer) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Storage.CookieState = store
	}
}

func SessionStorer(store abclientstate.SessionStorer) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Storage.SessionState = store
	}
}

func PreserveFields(fields ...string) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Modules.RegisterPreserveFields = fields
	}
}

func MountPath(path string) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Paths.Mount = path
	}
}

func RootURL(url string) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Paths.RootURL = url
	}
}

func ViewRenderer(renderer authboss.Renderer) ABOption {
	return func(ab *authboss.Authboss) {
		ab.Config.Core.ViewRenderer = renderer
	}
}

func SetCoreDefaults(ab *authboss.Authboss) {
	defaults.SetCore(&ab.Config, false, false)
}

func CoerceRedirectTo200(ab *authboss.Authboss) {
	r, ok := ab.Config.Core.Redirector.(*defaults.Redirector)
	if !ok {
		return
	}

	r.CorceRedirectTo200 = true
}

func Init(ab *authboss.Authboss) {
	if err := ab.Init(); err != nil {
		panic(err)
	}
}
