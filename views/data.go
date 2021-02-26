package views

const (
	AlertColorError   = "red"
	AlertColorSuccess = "green"
	AlertColorWarning = "orange"
	AlertColorInfo    = "blue"

	DefaultErrorTitle   = "Woops!"
	DefaultErrorMessage = "Something went wrong, please contact the admins!"
)

var DefaultErrorAlert = &Alert{
	Color:   AlertColorError,
	Title:   DefaultErrorTitle,
	Message: DefaultErrorMessage,
}

type Alert struct {
	Color   string
	Title   string
	Message string
}
type Data struct {
	Alert *Alert
	Yield interface{}
}

func (d *Data) SetAlert(err error)  {
	if pErr, ok := err.(PublicError); ok {
		d.Alert = &Alert{
			Color:   AlertColorError,
			Title:  DefaultErrorTitle,
			Message: pErr.Public(),
		}
	}else{
		d.Alert = DefaultErrorAlert
	}
}

func (d *Data) SetAlertError(message string)  {
	alert := DefaultErrorAlert
	alert.Message = message
	d.Alert = alert
}

type PublicError interface {
	error
	Public() string
}