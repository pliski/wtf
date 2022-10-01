package urlcheck

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wtfutil/wtf/logger"
)

type fUC func(uriResult) (StatusMsg, ErrMsg) // function Url Check

func CheckSomeUrl() fUC {
	return func(u uriResult) (StatusMsg, ErrMsg) {
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get(u.setting)

		defer func() {
			e := res.Body.Close()
			if e != nil {
				logger.Log(fmt.Sprintf("[urlcheck] response body close error: %s", e.Error()))
			}
		}()

		if err != nil {
			return 0, ErrMsg{err, true}
		}
		return StatusMsg(res.StatusCode), ErrMsg{nil, false}
	}
}

type StatusMsg int

type ErrMsg struct {
	err      error
	Critical bool
}

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e ErrMsg) Error() string { return e.err.Error() }
