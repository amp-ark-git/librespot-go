package respot

import (
	"github.com/art-media-platform/amp-librespot-go/librespot/core/crypto"
	"github.com/art-media-platform/amp-librespot-go/librespot/mercury"
	"github.com/art-media-platform/amp.SDK/stdlib/media"
	"github.com/art-media-platform/amp.SDK/stdlib/task"
)

// Forward declared method to create a new Spotify session
var StartNewSession func(ctx *SessionContext) (Session, error)

func DefaultSessionContext(deviceLabel string) *SessionContext {
	ctx := &SessionContext{
		DeviceName: deviceLabel,
	}
	return ctx
}

type SessionContext struct {
	task.Context              // logging & shutdown
	Login        SessionLogin // means for the session to login
	Info         SessionInfo  // filled in during Session.Login()
	Keys         crypto.Keys  // If left nil, will be auto-generated
	DeviceName   string       // Label of the device being used
	DeviceUID    string       // if nil, auto-generated from DeviceName
}

type SessionLogin struct {
	Username  string
	Password  string // AUTHENTICATION_USER_PASS
	AuthData  []byte // AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS
	AuthToken string // AUTHENTICATION_SPOTIFY_TOKEN
}

type SessionInfo struct {
	Username string // authenticated canonical username
	AuthBlob []byte // reusable authentication blob for Spotify Connect devices
	Country  string // user country returned by Spotify
}

type Session interface {
	Close() error

	// Returns the SessionContext current in use by this session
	Context() *SessionContext

	// Initiates login with params contained in Ctx.Login
	Login() error

	Search(query string, limit int) (*mercury.SearchResponse, error)
	Mercury() *mercury.Client

	// Initiates access ("pinning") with the given spotify track ID or URI
	PinTrack(trackID string, opts PinOpts) (media.Asset, error)
}

type PinOpts struct {

	// If set, media.Asset.OnStart(Ctx().Context) will be called on the returned media.Asset.
	// This is for convenience but not desirable when the asset is in a time-to-live cache, for example.
	StartInternally bool
}
