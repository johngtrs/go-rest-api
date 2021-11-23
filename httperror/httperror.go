package httperror

import (
	"errors"
	"net/http"
)

// 4**
var (
	ErrBadRequest                   = errors.New(http.StatusText(http.StatusBadRequest))                   // 400
	ErrUnauthorized                 = errors.New(http.StatusText(http.StatusUnauthorized))                 // 401
	ErrPaymentRequired              = errors.New(http.StatusText(http.StatusPaymentRequired))              // 402
	ErrForbidden                    = errors.New(http.StatusText(http.StatusForbidden))                    // 403
	ErrNotFound                     = errors.New(http.StatusText(http.StatusNotFound))                     // 404
	ErrMethodNotAllowed             = errors.New(http.StatusText(http.StatusMethodNotAllowed))             // 405
	ErrNotAcceptable                = errors.New(http.StatusText(http.StatusNotAcceptable))                // 406
	ErrProxyAuthRequired            = errors.New(http.StatusText(http.StatusProxyAuthRequired))            // 407
	ErrRequestTimeout               = errors.New(http.StatusText(http.StatusRequestTimeout))               // 408
	ErrConflict                     = errors.New(http.StatusText(http.StatusConflict))                     // 409
	ErrGone                         = errors.New(http.StatusText(http.StatusGone))                         // 410
	ErrLengthRequired               = errors.New(http.StatusText(http.StatusLengthRequired))               // 411
	ErrPreconditionFailed           = errors.New(http.StatusText(http.StatusPreconditionFailed))           // 412
	ErrRequestEntityTooLarge        = errors.New(http.StatusText(http.StatusRequestEntityTooLarge))        // 413
	ErrRequestURITooLong            = errors.New(http.StatusText(http.StatusRequestURITooLong))            // 414
	ErrUnsupportedMediaType         = errors.New(http.StatusText(http.StatusUnsupportedMediaType))         // 415
	ErrRequestedRangeNotSatisfiable = errors.New(http.StatusText(http.StatusRequestedRangeNotSatisfiable)) // 416
	ErrExpectationFailed            = errors.New(http.StatusText(http.StatusExpectationFailed))            // 417
	ErrTeapot                       = errors.New(http.StatusText(http.StatusTeapot))                       // 418
	ErrMisdirectedRequest           = errors.New(http.StatusText(http.StatusMisdirectedRequest))           // 421
	ErrUnprocessableEntity          = errors.New(http.StatusText(http.StatusUnprocessableEntity))          // 422
	ErrLocked                       = errors.New(http.StatusText(http.StatusLocked))                       // 423
	ErrFailedDependency             = errors.New(http.StatusText(http.StatusFailedDependency))             // 424
	ErrTooEarly                     = errors.New(http.StatusText(http.StatusTooEarly))                     // 425
	ErrUpgradeRequired              = errors.New(http.StatusText(http.StatusUpgradeRequired))              // 426
	ErrPreconditionRequired         = errors.New(http.StatusText(http.StatusPreconditionRequired))         // 428
	ErrTooManyRequests              = errors.New(http.StatusText(http.StatusTooManyRequests))              // 429
	ErrRequestHeaderFieldsTooLarge  = errors.New(http.StatusText(http.StatusRequestHeaderFieldsTooLarge))  // 431
	ErrUnavailableForLegalReasons   = errors.New(http.StatusText(http.StatusUnavailableForLegalReasons))   // 451

	// 5**
	ErrInternalServerError           = errors.New(http.StatusText(http.StatusInternalServerError))           // 500
	ErrNotImplemented                = errors.New(http.StatusText(http.StatusNotImplemented))                // 501
	ErrBadGateway                    = errors.New(http.StatusText(http.StatusBadGateway))                    // 502
	ErrServiceUnavailable            = errors.New(http.StatusText(http.StatusServiceUnavailable))            // 503
	ErrGatewayTimeout                = errors.New(http.StatusText(http.StatusGatewayTimeout))                // 504
	ErrHTTPVersionNotSupported       = errors.New(http.StatusText(http.StatusHTTPVersionNotSupported))       // 505
	ErrVariantAlsoNegotiates         = errors.New(http.StatusText(http.StatusVariantAlsoNegotiates))         // 506
	ErrInsufficientStorage           = errors.New(http.StatusText(http.StatusInsufficientStorage))           // 507
	ErrLoopDetected                  = errors.New(http.StatusText(http.StatusLoopDetected))                  // 508
	ErrNotExtended                   = errors.New(http.StatusText(http.StatusNotExtended))                   // 510
	ErrNetworkAuthenticationRequired = errors.New(http.StatusText(http.StatusNetworkAuthenticationRequired)) // 511
)
