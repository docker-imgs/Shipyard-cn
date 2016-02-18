package audit

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/shipyard/shipyard"
	"github.com/shipyard/shipyard/controller/manager"
)

var (
	ErrNoUserInToken = errors.New("没有收到用户令牌")
)

type Auditor struct {
	manager  manager.Manager
	excludes []string
}

// parses username from auth token
func getAuthUsername(r *http.Request) (string, error) {
	authToken := r.Header.Get("X-Access-Token")

	parts := strings.Split(authToken, ":")

	if len(parts) != 2 {
		return "", ErrNoUserInToken
	}

	return parts[0], nil
}

func filterURI(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	return u.Path, nil
}

func NewAuditor(m manager.Manager, excludes []string) *Auditor {
	return &Auditor{
		manager:  m,
		excludes: excludes,
	}
}

func (a *Auditor) HandlerFuncWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	skipAudit := false

	user, err := getAuthUsername(r)
	if err != nil {
		log.Errorf("审计错误: %s", err)
	}

	path, err := filterURI(r.RequestURI)
	if err != nil {
		log.Errorf("审计路径过滤错误: %s", err)
	}

	// check if excluded
	for _, e := range a.excludes {
		match, err := regexp.MatchString(e, path)
		if err != nil {
			log.Errorf("审计排除错误: %s", err)
		}

		if match {
			skipAudit = true
			break
		}
	}

	if user != "" && path != "" && !skipAudit {
		tagParts := strings.Split(path, "/")
		tag := tagParts[1]

		evt := &shipyard.Event{
			Type:     "api",
			Time:     time.Now(),
			Username: user,
			Message:  path,
			Tags:     []string{"api", tag, strings.ToLower(r.Method)},
		}

		if err := a.manager.SaveEvent(evt); err != nil {
			log.Errorf("保存事件错误: %s", err)
		}
	}

	log.Debugf("%s: %s", r.Method, r.RequestURI)

	// next must be called or middleware chain will break
	if next != nil {
		next(w, r)
	}
}
