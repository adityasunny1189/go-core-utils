package models

type ErrorType string

const (
	BOTTOM_OVERLAY ErrorType = "BOTTOM_OVERLAY"
	SNACK_BAR      ErrorType = "SNACK_BAR"
	IN_SCREEN      ErrorType = "IN_SCREEN"
	POPUP          ErrorType = "POPUP"
	TOAST          ErrorType = "TOAST"
)

type ApiResponse struct {
	Status               int         `json:"status"`
	Message              string      `json:"message"`
	Timestamp            string      `json:"timestamp"`
	Data                 interface{} `json:"data,omitempty"`
	ErrorHandlingDetails interface{} `json:"errorHandlingDetails,omitempty"`
}

type ErrorHandlingDetails struct {
	ErrorType    ErrorType   `json:"errorType"`
	ErrorDetails interface{} `json:"errorDetails"`
	ErrorCode    string      `json:"errorCode,omitempty"`
	UserFacing   bool        `json:"userFacing,omitempty"`
}

type BottomOverlayErrorDetails struct {
	IsFullScreen   bool              `json:"isFullScreen"`
	ImageURL       string            `json:"imageUrl"`
	Heading        string            `json:"heading"`
	Content        string            `json:"content"`
	AdditionalInfo map[string]string `json:"additionalInfo,omitempty"`
}

type SnakBarErrorDetails struct {
	Message string `json:"message"`
}

type InScreenErrorDetails struct {
	Color           string `json:"color"`
	Message         string `json:"message"`
	BackgroundImage string `json:"backgroundImage,omitempty"`
	ButtonText      string `json:"buttonText,omitempty"`
}

type PopUpErrorDetails struct {
	Message string `json:"message"`
}

type ToastErrorDetails struct {
	ToastTimerInSec int    `json:"toastTimerInSec"`
	ToastColor      string `json:"toastColor"`
	Message         string `json:"message"`
	IconURL         string `json:"iconUrl,omitempty"`
}
