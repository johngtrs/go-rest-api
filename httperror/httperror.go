package httperror

import (
	"errors"
	"net/http"
)

// 4**
var ErrBadRequest = errors.New(http.StatusText(http.StatusBadRequest))                                     // 400
var ErrUnauthorized = errors.New(http.StatusText(http.StatusUnauthorized))                                 // 401
var ErrPaymentRequired = errors.New(http.StatusText(http.StatusPaymentRequired))                           // 402
var ErrForbidden = errors.New(http.StatusText(http.StatusForbidden))                                       // 403
var ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))                                         // 404
var ErrMethodNotAllowed = errors.New(http.StatusText(http.StatusMethodNotAllowed))                         // 405
var ErrNotAcceptable = errors.New(http.StatusText(http.StatusNotAcceptable))                               // 406
var ErrProxyAuthRequired = errors.New(http.StatusText(http.StatusProxyAuthRequired))                       // 407
var ErrRequestTimeout = errors.New(http.StatusText(http.StatusRequestTimeout))                             // 408
var ErrConflict = errors.New(http.StatusText(http.StatusConflict))                                         // 409
var ErrGone = errors.New(http.StatusText(http.StatusGone))                                                 // 410
var ErrLengthRequired = errors.New(http.StatusText(http.StatusLengthRequired))                             // 411
var ErrPreconditionFailed = errors.New(http.StatusText(http.StatusPreconditionFailed))                     // 412
var ErrRequestEntityTooLarge = errors.New(http.StatusText(http.StatusRequestEntityTooLarge))               // 413
var ErrRequestURITooLong = errors.New(http.StatusText(http.StatusRequestURITooLong))                       // 414
var ErrUnsupportedMediaType = errors.New(http.StatusText(http.StatusUnsupportedMediaType))                 // 415
var ErrRequestedRangeNotSatisfiable = errors.New(http.StatusText(http.StatusRequestedRangeNotSatisfiable)) // 416
var ErrExpectationFailed = errors.New(http.StatusText(http.StatusExpectationFailed))                       // 417
var ErrTeapot = errors.New(http.StatusText(http.StatusTeapot))                                             // 418
var ErrMisdirectedRequest = errors.New(http.StatusText(http.StatusMisdirectedRequest))                     // 421
var ErrUnprocessableEntity = errors.New(http.StatusText(http.StatusUnprocessableEntity))                   // 422
var ErrLocked = errors.New(http.StatusText(http.StatusLocked))                                             // 423
var ErrFailedDependency = errors.New(http.StatusText(http.StatusFailedDependency))                         // 424
var ErrTooEarly = errors.New(http.StatusText(http.StatusTooEarly))                                         // 425
var ErrUpgradeRequired = errors.New(http.StatusText(http.StatusUpgradeRequired))                           // 426
var ErrPreconditionRequired = errors.New(http.StatusText(http.StatusPreconditionRequired))                 // 428
var ErrTooManyRequests = errors.New(http.StatusText(http.StatusTooManyRequests))                           // 429
var ErrRequestHeaderFieldsTooLarge = errors.New(http.StatusText(http.StatusRequestHeaderFieldsTooLarge))   // 431
var ErrUnavailableForLegalReasons = errors.New(http.StatusText(http.StatusUnavailableForLegalReasons))     // 451

// 5**
var ErrInternalServerError = errors.New(http.StatusText(http.StatusInternalServerError))                     // 500
var ErrNotImplemented = errors.New(http.StatusText(http.StatusNotImplemented))                               // 501
var ErrBadGateway = errors.New(http.StatusText(http.StatusBadGateway))                                       // 502
var ErrServiceUnavailable = errors.New(http.StatusText(http.StatusServiceUnavailable))                       // 503
var ErrGatewayTimeout = errors.New(http.StatusText(http.StatusGatewayTimeout))                               // 504
var ErrHTTPVersionNotSupported = errors.New(http.StatusText(http.StatusHTTPVersionNotSupported))             // 505
var ErrVariantAlsoNegotiates = errors.New(http.StatusText(http.StatusVariantAlsoNegotiates))                 // 506
var ErrInsufficientStorage = errors.New(http.StatusText(http.StatusInsufficientStorage))                     // 507
var ErrLoopDetected = errors.New(http.StatusText(http.StatusLoopDetected))                                   // 508
var ErrNotExtended = errors.New(http.StatusText(http.StatusNotExtended))                                     // 510
var ErrNetworkAuthenticationRequired = errors.New(http.StatusText(http.StatusNetworkAuthenticationRequired)) // 511
